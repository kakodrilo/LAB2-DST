/*
	Cliente
	Programa que simula los dos  tipos de clientes
*/
package main

// Se importan las librerias a utilizar
import(
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	pb "github.com/kakodrilo/LAB2-DST/pb" // libreria que contiene los compilados de protobuf - grpc
	"google.golang.org/grpc"
	"context"
	"log"
	"math/rand"
	"io"
	"time"
)

// se define la estructura que almacena cada chunk del archivo a subir y a descargar
type Chunk struct{
	chunk_id string
	data []byte
}

// SeparateFile: Recibe un nombre de archivo que se divide en chunks de 250 KB
// parametros:
// - file_name : nombre del archivo a descomponer
// retornos:
// - []Chunk arreglo con los chunks que componen el archivo
func SeparateFile(file_name string) []Chunk{
	fmt.Println("Separando el archivo")
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var size int64 = fileInfo.Size()

	chunk_size := 250 * (1 << 10)

	num_part := uint64(math.Ceil(float64(size) / float64(chunk_size)))

	arreglo := make([]Chunk, num_part)


	for i := uint64(0); i < num_part; i++ {

		partSize := int(math.Min(float64(chunk_size), float64(size-int64(i*uint64(chunk_size)))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		arreglo[i] = Chunk{
			chunk_id : strings.ReplaceAll(file_name," ","_") + "#" + strconv.FormatUint(i+uint64(1), 10),
			data : partBuffer}
		
	}

	return arreglo
}

// JoinFile: Recibe un arreglo de Chunk y construye el archivo que representan
// parametros:
// - parts : arreglo con los chunks del archivo a construir
// - book_name : nombre que se le pondra al archivo
func JoinFile(parts []Chunk, book_name string) {
	fmt.Println("Uniendo el archivo")
	_, err := os.Create(book_name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(book_name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalPartsNum := uint64(len(parts))

	var writePosition int64 = 0

	for i := uint64(0); i < totalPartsNum; i++ {

		var chunkSize int64 = int64(len(parts[i].data)) 
		chunkBufferBytes := make([]byte, chunkSize)

		writePosition = writePosition + chunkSize

		chunkBufferBytes = parts[i].data

		_, err := file.Write(chunkBufferBytes)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		file.Sync()

		chunkBufferBytes = nil 
	}

	file.Close()
}

// RequestBooks : Solicita los libors disponibles al NameNode y retorba el arreglo con las direcciones ip de las partes del libro seleccionado
// retornos:
// - []string arreglo con las direcciones ip de las partes de la eleccion
// - string nombre del libro seleccionado
func RequestBooks() ([]string, string) {
	fmt.Println("Solicitando libros disponibles")
	connection, err := grpc.Dial("dist08:9008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con NameNode: %v", err)
	}
	
	defer connection.Close()

	client := pb.NewNameNodeClient(connection)

	books_streaming, err := client.FileRequest(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatalf("No se pueden solocitar los libros: %v", err)
	}

	var books []string

	for {
		book, err := books_streaming.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}

		books = append(books, book.FileName)
	}

	//connection.Close()

	fmt.Println(" Seleccione el libro que desea descargar: \n")
	for i := 1; i <= len(books); i++ {
		fmt.Printf(" [%d] %s \n", i, books[i-1])
	}

	fmt.Println("\n Ingrese opcion: ")

	var option int32
	fmt.Scan(&option)

	book_name := books[option-1]
	fmt.Println("Solicitando direcciones")
	res, err := client.AddressRequest(context.Background(), &pb.File{ FileName: book_name})

	if err != nil {
		log.Fatalf("No se pueden solocitar las direcciones de los Chunks: %v", err)
	}

	return res.Ip, book_name
}

// RequestChunks: Solicita los chunks en las direcciones entregadas por parametro
// parametros:
// - addresses : arreglo con las direcciones ip de las partes 
// - file_name : nombre del archivo
// retornos:
// - []Chunk arreglo con los chunks que componen el archivo
func RequestChunks(addresses []string, file_name string) []Chunk{
	fmt.Println("Solicitando partes")
	var array []Chunk;
	for i,ip := range addresses{
			con := ip+":900"+string(ip[10])
			connection, err := grpc.Dial(con, grpc.WithInsecure())

			if err != nil {
				log.Fatalf("Error al conectarse con DataNode: %v", err)
			}

			client := pb.NewDataNodeClient(connection)

			res, err := client.DownloadChunks(context.Background(), &pb.RequestChunk{ChunkId: file_name+"#"+strconv.Itoa(i+1)})

			if err != nil {
				log.Fatalf("Error al descargar chunks: %v", err)
			}

			array = append(array, Chunk{chunk_id: res.ChunkId , data: res.Chunk })

			connection.Close()
	}

	return array

}

// SendChunks : Envia los chunks a algun datanode activo seleccionado aleatoriamente
// parametros:
// - chunks : arreglo con los chunks a enviar
func SendChunks(chunks []Chunk) {
	fmt.Println("Subiendo chunks del archivo")
	fmt.Printf("Chunks a subir: %v \n", len(chunks))
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for {
		election := r1.Int31n(3)
		fmt.Printf("Se escogio el nodo " + "10.6.40.14"+strconv.Itoa(5+int(election))+":900"+strconv.Itoa(5+int(election)) + "\n" )
		connection, err := grpc.Dial("10.6.40.14"+strconv.Itoa(5+int(election))+":900"+strconv.Itoa(5+int(election)), grpc.WithInsecure())
		if err != nil {
			fmt.Printf("error al conectar: %v \n", err)
		}else{
			fmt.Println("Conectando ...")
			client := pb.NewDataNodeClient(connection)

			stream, err := client.UploadChunks(context.Background())

			if err != nil {
				log.Println("Error al enviar contactar al nodo")
			}else{

				for _, chunk := range chunks {

					if err := stream.Send(&pb.Chunk{ChunkId: chunk.chunk_id, Chunk: chunk.data }); err != nil {
						log.Fatalf("%v.Send(%v) = %v", stream, chunk, err)
					}
				}

				reply, err := stream.CloseAndRecv()
				if err != nil {
					log.Fatalf("Error al enviar los chunks")
				}
				fmt.Println(reply.Response)
				break
			}
		}
	}

	
    
}


func main(){

	fmt.Println("\n Seleccione el tipo de cliente que quiere simular: ")

	fmt.Println(" [1] Cliente Uploader                  [2] Cliente Downloader")

	var option int
	fmt.Scan(&option)

	if option == 1 {
		fmt.Println(" Ingrese el nombre del archivo que desea subir (incluya la extension):")
		fmt.Println(" * Recuerde que el archivo debe estar en la misma carpeta que este codigo")
		fmt.Println(" * ejemplo: 'Dracula-Stoker_Bram.epub' ")

		var bookToUpload string
		fmt.Scan(&bookToUpload)

		chunks := SeparateFile(bookToUpload)

		fmt.Printf("El archivo se dividio en %v chunks\n", len(chunks))

		SendChunks(chunks)

	}else if option == 2{
		ips, book_name := RequestBooks()

		chunks := RequestChunks(ips, book_name)

		JoinFile(chunks, book_name)

	}

}