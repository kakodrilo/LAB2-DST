/*
	NameNode
	Programa que simula al NameNode
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
    "strconv"
    pb "github.com/kakodrilo/LAB2-DST/pb"
    "google.golang.org/grpc"
    "context"
    "time"
    "sync"
) 

// se define la estructura que permite asegurar secciones criticas
type Protection struct {
    flag int32
    cond *sync.Cond
}

// estructura para el funcionamiento del servidor para clientes y  DataNodes
type NameNodeServer struct{
    pb.UnimplementedNameNodeServer
}

// VARIABLES GLOBALES

// ip de los otrod DataNodes
var ip1 string = "10.6.40.145:9005"
var ip2 string = "10.6.40.146:9006"
var ip3 string = "10.6.40.147:9007"

// se inician la estructura de proteccion
var protection *Protection = &Protection{
    flag: 1,
    cond: &sync.Cond{L: &sync.Mutex{}}}

// map [nombre de libro] = direcciones ip ordenadas
var books map[string][]string = make( map[string][]string)

var mux sync.Mutex

// LoadBooks: Carga los datos guardados en el archivo Log.txt, si este no existe lo crea
func LoadBooks() { 

    file, err := os.Open("Log.txt") 
  
    if err != nil { 
        _ ,_ = os.Create("Log.txt")
        fmt.Println("Se ha creado el archivo Log ...")
        return
    } 
    fmt.Println("Cargando Datos del Log ...")
  
    scanner := bufio.NewScanner(file) 
    scanner.Split(bufio.ScanLines) 
    var text []string 
  
    for scanner.Scan() { 
        text = append(text, scanner.Text()) 
    } 

    file.Close() 
  
    cant_lines := len(text)

    var chunks []string

    for i := 0; i < cant_lines; i++{ 
        
        splitResult := strings.Split(text[i], " ")

        name := splitResult[0]
        cant_chunk, _ := strconv.Atoi(splitResult[1])

        var j int

        for j = i + 1; j < i + 1 + cant_chunk; j++{
            aux := strings.Split(text[j], " ")
            ip := aux[1]
            chunks = append(chunks, ip)
        } 

        books[name] = make([]string, cant_chunk)
        copy(books[name],chunks)
        //fmt.Println(books)

        i = j - 1

        //fmt.Println(name)
        //fmt.Printf("%v \n", chunks)

        chunks = chunks[:0]
    } 
} 

// FileRequest: Entrega los libros disponibles en respuesta a una solicitud de un cliente
// parametros:
// - empty : Mensaje utilizado para llamar a la funcion
// - stream : Requerida para enviar un conjunto de mensajes de tipo pb.File
// retornos:
// - pb.File mensaje con el nombre de uno de los libros disponibles
func (s *NameNodeServer) FileRequest(empty *pb.Empty, stream pb.NameNode_FileRequestServer) error {
    mux.Lock()
    fmt.Println("Recibiendop solicitud de archivos")
    for book := range books {
        file := &pb.File{FileName: book}
        err := stream.Send(file)
        if err != nil {
            mux.Unlock()
            log.Fatalf("failed to response FileRequest: %v", err) 
        }
    }
    mux.Unlock()
    return nil
}

// AddressRequest: Entrega las direcciones ip de las maquinas que almacenan los chunks del libro solicitado
// parametros:
// - ctx : variable necesaria para la immplementacion de funciones en protobuf
// - file : Mensaje con le nombre del libro requerido
// retornos:
// - pb.ChunkAddres mensaje que contiene el nombre del libro solicitado y un arreglo con las ip donde se encuentran las partes
func (s *NameNodeServer) AddressRequest( ctx context.Context, file *pb.File) ( *pb.ChunkAddress,error) {
    return &pb.ChunkAddress{Ip: books[file.FileName]}, nil
}

// Server: se da soporte al server para clientes y DataNodes mediante protocol buffers - grpc
func Server(){
    fmt.Println("Iniciando Server")
	lis, err := net.Listen("tcp", ":9008")
	if err != nil {
		log.Fatalf("No se pudo iniciar el server: %v", err)
	}
	// se inicia el server
	s := grpc.NewServer()
	// se registra el server pora dar soporte al servicio de clientes
	pb.RegisterNameNodeServer(s , &NameNodeServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("No se pudo crear el servidor: %v", err)
	}
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

// FinalProposal: Recibe la propuesta de un DaataNode con algoritmo centralizado, se la envia a los DataNodes y la modifica si es necesario
// parametros:
// - ctx : variable necesaria para la immplementacion de funciones en protobuf
// - proposal : mensaje que contiene el nombre del libro y las partes que guardara cada nodo
// retornos:
// - pb.Proposal propuesta de respuesta que puede estar modificada si corresponde
func (s *NameNodeServer) FinalProposal( ctx context.Context, proposal *pb.Proposal) ( *pb.Proposal,error) {
    fmt.Println("Recibiendo proposicion")
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

    if len(reasign) == 0 || ((flag1==1 || len(proposal.ChunksDn1)==0) && (flag2==1 || len(proposal.ChunksDn2)==0) && (flag3==1 || len(proposal.ChunksDn3)==0)){
        return proposal, nil
    }else{
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
        return newProposal,nil

    }

}

// WriteFile: Escribe en el archivo Log el nuevo libro ingresado, y agrega la informacion al map global
// parametros:
// - file_name : nombre del libro a guardar
// - dn1 : arreglo de enteros con las partes que guardara el DataNode1
// - dn2 : arreglo de enteros con las partes que guardara el DataNode2
// - dn3 : arreglo de enteros con las partes que guardara el DataNode3
// retornos:
// - error para saber si la escritura falla o no 
func WriteFile(file_name string, dn1 []int32, dn2 []int32, dn3 []int32) error{
    fmt.Println("Escribiendo Log")
    positions := make([]int32, len(dn1)+len(dn2)+len(dn3))
    
    max := len(dn1)
    if max < len(dn2){
        max = len(dn2)
    }
    if max < len(dn3){
        max = len(dn3)
    }

    for i := 0; i < max; i++ {
        if i < len(dn1) {
            positions[dn1[i]-int32(1)] = 1
        } 
        if i < len(dn2) {
            positions[dn2[i]-int32(1)] = 2
        }
        if i < len(dn3) {
            positions[dn3[i]-int32(1)] = 3
        }
    }
    // acquire candado
    mux.Lock()
    file, err := os.OpenFile("Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
        log.Printf("Fallo al abrir el archivo: %v", err)
        mux.Unlock()  
        return err
    }
    _,err = file.WriteString(file_name + " " + strconv.Itoa(len(positions)) + "\n")
    if err != nil {
        log.Printf("Fallo al escribir en el archivo: %v", err)
        mux.Unlock()  
        return err
    }
    books[file_name] = []string{}
    for i, dn := range positions {
        _,err := file.WriteString("parte_" + strconv.Itoa(i+1) + " 10.6.40.14" + strconv.Itoa(4+int(dn)) + "\n")
        if err != nil {
            log.Printf("Fallo al escribir en el archivo: %v", err)
            mux.Unlock()  
            return err 
        }

        books[file_name] = append(books[file_name], "10.6.40.14" + strconv.Itoa(4+int(dn)))
    }
    file.Close()

    // release candado
    mux.Unlock()    
    return nil
}

// RequestWrite: Solicitud de escritura en el archivo
// parametros:
// - ctx : variable necesaria para la immplementacion de funciones en protobuf
// - empty : Mensaje utilizado para llamar a la funcion
// retornos:
// - pb.Response respuesta a la operacion realizada
// - error para saber si la operacion falla o no 
func (s *NameNodeServer) RequestWrite( ctx context.Context, proposal *pb.Empty) ( *pb.Response,error) {

    protection.cond.L.Lock()
    fmt.Println("Recibiendo solicitud de escritura")
    for protection.flag == 0 {
        protection.cond.Wait()
    }
    protection.flag = 0
    protection.cond.L.Unlock()

    return &pb.Response{Response: "ok"},nil

}

// CentralizedWriteLog: da soporte a la escritura del archivo por peticion de un DataNode con algoritmo centralizado
// parametros:
// - ctx : variable necesaria para la immplementacion de funciones en protobuf
// - proposal : mensaje que contiene el nombre del libro y las partes que guardara cada nodo paara ser escritos
// retornos:
// - pb.Response respuesta a la operacion realizada
// - error para saber si la operacion falla o no 
func (s *NameNodeServer) CentralizedWriteLog( ctx context.Context, proposal *pb.Proposal) ( *pb.Response,error) {

    protection.cond.L.Lock()
    err := WriteFile(proposal.FileName, proposal.ChunksDn1, proposal.ChunksDn2, proposal.ChunksDn3)
    protection.flag = 1
    protection.cond.L.Unlock()
    protection.cond.Signal()
    if err != nil {
        return &pb.Response{Response: fmt.Sprintf("error: %v",err)}, nil
    }


    return &pb.Response{Response: "ok"}, nil
}

// DistributedWriteLog: da soporte a la escritura del archivo por peticion de un DataNode con algoritmo distribuido
// parametros:
// - ctx : variable necesaria para la immplementacion de funciones en protobuf
// - proposal : mensaje que contiene el nombre del libro y las partes que guardara cada nodo paara ser escritos
// retornos:
// - pb.Response respuesta a la operacion realizada
// - error para saber si la operacion falla o no 
func (s *NameNodeServer) DistributedWriteLog( ctx context.Context, proposal *pb.Proposal) ( *pb.Response,error) {

    err := WriteFile(proposal.FileName, proposal.ChunksDn1, proposal.ChunksDn2, proposal.ChunksDn3)
    if err != nil {
        return &pb.Response{Response: fmt.Sprintf("error: %v",err)}, nil
    }
    return &pb.Response{Response: "ok"}, nil
}

func main(){
    LoadBooks()
    Server()
}