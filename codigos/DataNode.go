/*
	DataNode
	Programa que simula a los DataNodes con algoritmo distribuido o centralizado
*/
package main

// se importan la librerias necesarias para el funcionamiento del codigo
import ( 
    "net"
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    pb "github.com/kakodrilo/LAB2-DST/pb" // libreria que contiene los compilasdos de protobuf - grpc
    "google.golang.org/grpc"
    "context"
    "time"
	"sync"
	"io/ioutil"
	"io"
) 

// estructura para el funcionamiento del servidor para clientes y otros DataNodes
type DataNodeServer struct{
    pb.UnimplementedDataNodeServer
}

// se define la estructura que representa a los chunks a almacenar y distribuir
type Chunk struct{
	chunk_id string
	data []byte
}

// se define la estructura que permite asegurar secciones criticas
type Protection struct{
	timeStamp string
	state int32
	cond *sync.Cond
}

// VARIABLES GLOBALES

// se inician la estructura de proteccion
var protection *Protection = &Protection{
	timeStamp: "",
	state: 1 ,
	cond: &sync.Cond{L: &sync.Mutex{}}}
	

// ip de los otrod DataNodes
var ip1 string = "10.6.40.145:9005"
var ip2 string = "10.6.40.146:9006"
var ip3 string = "10.6.40.147:9007"
// identificador del Nodo
var NodeId int = 3


var option int32 // 1 es distribuido    0 es centralizado

var timeStampFormat string = "Jan _2 15:04:05.000000"

var mux sync.Mutex


// UploadChunks: Recibe los chunks desde el cliente y los distribuye segun la opcion elegida
// parametros:
// - stream : Requerida para recibir un conjunto de mensajes de tipo pb.Chunk
func (s *DataNodeServer) UploadChunks(stream pb.DataNode_UploadChunksServer) error {
	fmt.Println("Recibiendo Chunks")
	var chunks []Chunk

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunks = append(chunks, Chunk{chunk_id: res.ChunkId , data: res.Chunk })
	}
	mux.Lock()
	if option == 2 {
		DistributedUpload(chunks)
	}else if option == 1{
		CentralizedUpload(chunks)
	}
	mux.Unlock()
	return stream.SendAndClose(&pb.Response{Response: "Libro subido y distribuido correctamente."})

}

// SeparateFile: Recibe una proposicion de acceso al recurso que trae un timestamp para comparar y quien realizo la proposicion
// parametros:
// - ctx : requerida para la implememtacion de bunciones protobuf
// - proposal : proposicion de acceso
// retornos:
// - pb.Response respuesta a la solicitud
func (s *DataNodeServer) ProposalRequest(ctx context.Context, proposal *pb.ProposalAcces) (*pb.Response, error){
	
	protection.cond.L.Lock()
	fmt.Printf("Recibiendo proposicion de %v \n", proposal.NodeId)
	var myTime time.Time
	if protection.state == 2 {
		myTime, _  = time.Parse(timeStampFormat, protection.timeStamp)
	}

	requesTime, _ := time.Parse(timeStampFormat, proposal.TimeStamp)

	// esperar hasta que este libre para responder
    for (protection.state == 0) || ((protection.state == 2) && myTime.Before(requesTime)) {
		protection.cond.Wait()
		if protection.state == 2 {
			myTime, _  = time.Parse(timeStampFormat, protection.timeStamp)
		}
    }
    protection.cond.L.Unlock()
	// mandar ok

	return &pb.Response{Response: "ok"},nil
}

// ping: verifica que la maquina con cierta ip este disponible
// parametros:
// - ip : atring con la ip de la maquina a consultar
// retornos:
// - bool  true si esta disponible, false si no
func ping(ip string) bool {
	fmt.Println("Haciendo ping a " + ip)
    Time := time.Duration(3 * time.Second)
    _, err := net.DialTimeout("tcp", ip, Time)
    if err != nil{
      return false
    }
    return true
}

// DistributedUpload:Implementa el algoritmo distribuido
// parametros:
// - chunks : arreglo de chunks a repartir
func DistributedUpload(chunks []Chunk){
	fmt.Println("Iniciando Subida Distribuida")
	// saber que nodos estan disponibles
	
	proposal := &pb.Proposal{}
	fmt.Printf("Cantidad de chunks a repartir: %v \n",len(chunks))
	splitResult := strings.Split(chunks[0].chunk_id , "#")
	proposal.FileName = splitResult[0]

	intDiv := len(chunks)/3

	if len(chunks) == 1 {
		proposal.ChunksDn1 =[]int32{1}
		proposal.ChunksDn2 =[]int32{}
		proposal.ChunksDn3 =[]int32{}	 
	}else if len(chunks) == 2 {
		proposal.ChunksDn1 =[]int32{1}
		proposal.ChunksDn2 =[]int32{2}
		proposal.ChunksDn3 =[]int32{}	
	}else if len(chunks) >= 3 {
		proposal.ChunksDn1 = []int32{}
		proposal.ChunksDn2 = []int32{}
		proposal.ChunksDn3 = []int32{}
		for i := 1; i <= len(chunks); i++ {
			if i <= intDiv {
				proposal.ChunksDn1 = append(proposal.ChunksDn1, int32(i))
			}else if i <= 2*intDiv {
				proposal.ChunksDn2 = append(proposal.ChunksDn2, int32(i))
			}else {
				proposal.ChunksDn3 = append(proposal.ChunksDn3, int32(i))
			}
		}
	}

	var newProposal *pb.Proposal = &pb.Proposal{}
    newProposal.FileName = proposal.FileName
    flag1 := 1
    flag2 := 1
    flag3 := 1

    var reasign []int32

    if !ping(ip1){
        flag1 = 0
        reasign = append(reasign, proposal.ChunksDn1...)
        newProposal.ChunksDn1 = []int32{}
    }else{
        newProposal.ChunksDn1 = proposal.ChunksDn1
    } 

    if !ping(ip2){
        flag2 = 0
        reasign = append(reasign, proposal.ChunksDn2...)
        newProposal.ChunksDn2 = []int32{}
    }else{
        newProposal.ChunksDn2 = proposal.ChunksDn2
    } 

    if !ping(ip3){
        flag3 = 0
        reasign = append(reasign, proposal.ChunksDn3...)
        newProposal.ChunksDn3 = []int32{}
    }else{
        newProposal.ChunksDn3 = proposal.ChunksDn3
    } 
	
	if !(len(reasign) == 0 || ((flag1==1 || len(proposal.ChunksDn1)==0) && (flag2==1 || len(proposal.ChunksDn2)==0) && (flag3==1 || len(proposal.ChunksDn3)==0))){
        fmt.Println("Cambiando proposicion")
        actives := flag1+flag2+flag3

        if actives == 2 {
            if len(reasign) == 1 {
                if flag1 == 1 && len(proposal.ChunksDn1)==0 {
                    newProposal.ChunksDn1 = append(proposal.ChunksDn1, reasign...)
                }else if flag2 == 1 && len(proposal.ChunksDn2)==0 {
                    newProposal.ChunksDn2 = append(proposal.ChunksDn2, reasign...)
                }else if flag3 == 1 && len(proposal.ChunksDn3)==0 {
                    newProposal.ChunksDn3 = append(proposal.ChunksDn3, reasign...)
                }else{
                    if flag1 == 1  {
                        newProposal.ChunksDn1 = append(proposal.ChunksDn1, reasign...)
                    }else if flag2 == 1{
                        newProposal.ChunksDn2 = append(proposal.ChunksDn2, reasign...)
                    }else if flag3 == 1{
                        newProposal.ChunksDn3 = append(proposal.ChunksDn3, reasign...)
                    }
                }
            }else{
                if flag1 == 1 && flag2 == 1{
                    newProposal.ChunksDn1 = append(proposal.ChunksDn1, reasign[:len(reasign)/2]...)
                    newProposal.ChunksDn2 = append(proposal.ChunksDn2, reasign[len(reasign)/2:]...)
                }else if flag1 == 1 && flag3 == 1{
                    newProposal.ChunksDn1 = append(proposal.ChunksDn1, reasign[:len(reasign)/2]...)
                    newProposal.ChunksDn3 = append(proposal.ChunksDn3, reasign[len(reasign)/2:]...)
                }else if flag3 == 1 && flag2 == 1{
                    newProposal.ChunksDn3 = append(proposal.ChunksDn3, reasign[:len(reasign)/2]...)
                    newProposal.ChunksDn2 = append(proposal.ChunksDn2, reasign[len(reasign)/2:]...)
                }
            }

        }else if actives == 1{
            if flag1 == 1 {
                newProposal.ChunksDn1 = append(proposal.ChunksDn1, reasign...)
            }else if flag2 == 1 {
                newProposal.ChunksDn2 = append(proposal.ChunksDn2, reasign...)
            }else if flag3 == 1 {
                newProposal.ChunksDn3 = append(proposal.ChunksDn3, reasign...)
            }
        }
        proposal = newProposal

	}
	
	if NodeId == 1 {
		LocalChunks(ExtractChunks(chunks, proposal.ChunksDn1))
		if len(proposal.ChunksDn2) > 0 {
			SendChunks(ip2, ExtractChunks(chunks, proposal.ChunksDn2))
		}
		if len(proposal.ChunksDn3) > 0 {
			SendChunks(ip3, ExtractChunks(chunks, proposal.ChunksDn3))
		}
	}else if NodeId == 2 {
		LocalChunks(ExtractChunks(chunks, proposal.ChunksDn2))
		if len(proposal.ChunksDn1) > 0 {
			SendChunks(ip1, ExtractChunks(chunks, proposal.ChunksDn1))
		}
		if len(proposal.ChunksDn3) > 0 {
			SendChunks(ip3, ExtractChunks(chunks, proposal.ChunksDn3))
		}
	}else{
		LocalChunks(ExtractChunks(chunks, proposal.ChunksDn3))
		if len(proposal.ChunksDn1) > 0 {
			SendChunks(ip1, ExtractChunks(chunks, proposal.ChunksDn1))
		}
		if len(proposal.ChunksDn2) > 0 {
			SendChunks(ip2, ExtractChunks(chunks, proposal.ChunksDn2))
		}
	}


	// paso a estado de solicitud de recurso y actualizo mi timestamp
	protection.cond.L.Lock()
	protection.state = 2
	current := time.Now() 
	protection.timeStamp = current.Format(timeStampFormat)
	protection.cond.L.Unlock()
	//protection.cond.Signal()

	// solicitar permiso al recurso
	// esperar a que todos respondan - tener cuidado con si hay nodos caidos (verificar lista de avaible)
	fmt.Println("Solicitando permiso para escribir")
	if flag1 == 1 {
		RequestAcces(ip1)
	}
	if flag2 == 1 {
		RequestAcces(ip2)
	}
	if flag3 == 1 {
		RequestAcces(ip3)
	}
	// Acceder al recurso
	protection.cond.L.Lock()
	fmt.Println("Escribiendo en el Log")
	protection.state = 0
	connection, err := grpc.Dial("10.6.40.148:9008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con NameNode: %v", err)
	}
	defer connection.Close()

	client := pb.NewNameNodeClient(connection)

	res, err := client.DistributedWriteLog(context.Background(), proposal)
	if err != nil {
		log.Fatalf("Error al registrar %s: %v", proposal.FileName, err)
	}
	protection.state = 1
	protection.cond.L.Unlock()
	protection.cond.Signal()

	fmt.Printf("La escritura finalizo con el estado: %s \n", res.Response)

}

// RequestAcces: Envia una solicitud de acceso al recurso compartido a la ip entregada
// parametros:
// - ip : string con la ip de la maquina a consultar
func RequestAcces(ip string) {
	fmt.Printf("Solicitando permiso a %v \n", ip)
	connection, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con %s: %v", ip, err)
	}
	defer connection.Close()
	
	client := pb.NewDataNodeClient(connection)

	res, err := client.ProposalRequest(context.Background(), &pb.ProposalAcces{NodeId: int32(NodeId), TimeStamp: protection.timeStamp})
	if err != nil {
		log.Fatalf("No se puedo enviar solicitud de escritura a %s: %v",ip, err)
	}

	if res.Response != "ok"{
		log.Fatalf("solicitud de escritura negativa de %s: %v",ip, err)
	}
}

// CentralizedUpload: Implemeta el algoritmo centralizado
// parametros:
// - chunks : arreglo con los chunks a repartir
func CentralizedUpload(chunks []Chunk){
	fmt.Println("Iniciando Subida Centralizada")
	proposal := &pb.Proposal{}
	fmt.Printf("Cantidad de chunks a repartir: %v \n",len(chunks))
	splitResult := strings.Split(chunks[0].chunk_id , "#")
	proposal.FileName = splitResult[0]

	intDiv := len(chunks)/3

	if len(chunks) == 1 {
		proposal.ChunksDn1 =[]int32{1}
		proposal.ChunksDn2 =[]int32{}
		proposal.ChunksDn3 =[]int32{}	 
	}else if len(chunks) == 2 {
		proposal.ChunksDn1 =[]int32{1}
		proposal.ChunksDn2 =[]int32{2}
		proposal.ChunksDn3 =[]int32{}	
	}else if len(chunks) >= 3 {
		proposal.ChunksDn1 = []int32{}
		proposal.ChunksDn2 = []int32{}
		proposal.ChunksDn3 = []int32{}
		for i := 1; i <= len(chunks); i++ {
			if i <= intDiv {
				proposal.ChunksDn1 = append(proposal.ChunksDn1, int32(i))
			}else if i <= 2*intDiv {
				proposal.ChunksDn2 = append(proposal.ChunksDn2, int32(i))
			}else {
				proposal.ChunksDn3 = append(proposal.ChunksDn3, int32(i))
			}
		}
	}

	fmt.Println("Enviando propuesta de distribucion a NameNode")
	connection, err := grpc.Dial("10.6.40.148:9008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con NameNode: %v", err)
	}
	defer connection.Close()

	client := pb.NewNameNodeClient(connection)

	res, err := client.FinalProposal(context.Background(), proposal)
	if err != nil {
		log.Fatalf("No se puede enviar la Proposion: %v", err)
	}

	//connection.Close()

	if NodeId == 1 {
		LocalChunks(ExtractChunks(chunks, res.ChunksDn1))
		if len(res.ChunksDn2) > 0 {
			SendChunks(ip2, ExtractChunks(chunks, res.ChunksDn2))
		}
		if len(res.ChunksDn3) > 0 {
			SendChunks(ip3, ExtractChunks(chunks, res.ChunksDn3))
		}
	}else if NodeId == 2 {
		LocalChunks(ExtractChunks(chunks, res.ChunksDn2))
		if len(res.ChunksDn1) > 0 {
			SendChunks(ip1, ExtractChunks(chunks, res.ChunksDn1))
		}
		if len(res.ChunksDn3) > 0 {
			SendChunks(ip3, ExtractChunks(chunks, res.ChunksDn3))
		}
	}else{
		LocalChunks(ExtractChunks(chunks, res.ChunksDn3))
		if len(res.ChunksDn1) > 0 {
			SendChunks(ip1, ExtractChunks(chunks, res.ChunksDn1))
		}
		if len(res.ChunksDn2) > 0 {
			SendChunks(ip2, ExtractChunks(chunks, res.ChunksDn2))
		}
	}

	if !(len(res.ChunksDn1) == 0 &&  len(res.ChunksDn2) == 0 && len(res.ChunksDn3) == 0){
		fmt.Println("Escribiendo Log")
		writeResponse, err := client.CentralizedWriteLog(context.Background(), res)
		if err != nil {
			log.Fatalf("No se puede escribir el en Log: %v", err)
		}
		fmt.Printf("La escritura finalizo con el estado: %s \n", writeResponse.Response)
	}
	
}


// ExtractChunks: retorna los chunks correpondientes a un arreglo de id 
// parametros:
// - chunks : arreglo con todo los chunks
// - chunksId : arreglo con las id que se desean
// retornos:
// - []Chunk arreglo con los chunks deseados
func ExtractChunks(chunks []Chunk, chunksId []int32) ([]Chunk){

	chunksResponse := []Chunk{}

	for _, id := range chunksId{
		chunksResponse = append(chunksResponse, chunks[id-1])
	}
	return chunksResponse

}

// LocalChunks: guarda de manera local los chunks entregados
// parametros:
// - chunks : arreglo con los chunks a guardar en disco
func LocalChunks(chunks []Chunk){
	fmt.Println("Guardando chunks locales")
	for _, chunk := range chunks {

		fileName := chunk.chunk_id
		_, err := os.Create("Libros/"+fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ioutil.WriteFile("Libros/"+fileName, chunk.data, os.ModeAppend)

		fmt.Println("Se ha guardado: ", fileName)

	}

}

// SendChunks : Envia los chunks entregados a la maquina de ip entregada
// parametros:
// - ip : ip y puerto de la aplicacion que recibe los chunks
// - chunks : arreglo de chunks a enviar
func SendChunks(ip string, chunks []Chunk){
	fmt.Printf("Enviando chunks a %v \n", ip)

	connection, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectarse con %s: %v \n", ip, err)
		return
	}
	defer connection.Close()

	client := pb.NewDataNodeClient(connection)

	stream, err := client.SaveChunks(context.Background())
	if err != nil {
		log.Printf("%v.RecordRoute(_) = _, %v \n", client, err)
		return
	}

	for _, chunk := range chunks {

		if err := stream.Send(&pb.Chunk{ChunkId: chunk.chunk_id, Chunk: chunk.data }); err != nil {
			log.Printf("%v.Send(%v) = %v \n", stream, chunk, err)
			return
		}
	}

	reply, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error al enviar los chunks \n")
		return
	}
	fmt.Println(reply.Response)
}

// SaveChunks: Recibe los chunks enviados por otro DataNode para ser almacenados localmente
// parametros:
// - stream : Requerida para recibir un conjunto de mensajes de tipo pb.Chunk
// retornos:
// - pb.Response respuesta a la operacion realizada
func (s *DataNodeServer) SaveChunks(stream pb.DataNode_SaveChunksServer) error {
	fmt.Println("Recibiendo chunks para guardar")
	var chunks []Chunk

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunks = append(chunks, Chunk{chunk_id: res.ChunkId , data: res.Chunk })
	}

	LocalChunks(chunks)

	return stream.SendAndClose(&pb.Response{Response: "Chunks almacenados correctamente."})
}

// DownloadChunks: entrega el chunk requerido por un cliente
// parametros:
// - ctx : variable requerida para la implementacion de funciones probuf
// - request : mensaje con la identificacion del chunk requerido
// retornos:
// - pb.Chunk Chunk solicitado
func (s *DataNodeServer) DownloadChunks(ctx context.Context, request *pb.RequestChunk) (*pb.Chunk, error) {
	fmt.Println("Recibiendo peticion de descarga de chunks")
	chunkFile, err := os.Open("Libros/" + request.ChunkId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer chunkFile.Close()

	chunkInfo, err := chunkFile.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var chunkSize int64 = chunkInfo.Size() 
	chunkBufferBytes := make([]byte, chunkSize)

	reader := bufio.NewReader(chunkFile)
	_, err = reader.Read(chunkBufferBytes)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	return &pb.Chunk{ChunkId: request.ChunkId, Chunk: chunkBufferBytes}, nil          
}

// Server: se da soporte al server para clientes y DataNodes mediante protocol buffers - grpc
// parametros:
// - port : puerto en donde escuchara el server
func Server(port string){
	fmt.Println(" Iniciando Server ...")
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("No se pudo iniciar el server: %v", err)
	}
	// se inicia el server
	s := grpc.NewServer()
	// se registra el server pora dar soporte al servicio de clientes
	pb.RegisterDataNodeServer(s , &DataNodeServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("No se pudo crear el servidor: %v", err)
	}
}


func main(){

	//preguntar tipo de distribuci[on]
	fmt.Println("\n Seleccione el tipo de Algortimo que quiere simular: ")
	fmt.Println("  [1] Centralizado              [2] Distribuido")
	fmt.Println("\n Ingrese Opcion: ")
	fmt.Scan(&option)

	
	Server("9007")
}
