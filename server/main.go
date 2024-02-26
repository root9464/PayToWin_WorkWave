package main

import "root/core"

func main() {
	core.Core()
}

// server1 := grpc.NewServer()
// server2 := grpc.NewServer()

// // Регистрация сервисов на сервере
// pb.RegisterService1Server(server1, &service1Impl{})
// pb.RegisterService2Server(server2, &service2Impl{})

// // Создание слушателей на разных портах
// lis1, err := net.Listen("tcp", ":3000")
// if err != nil {
//     log.Fatalf("failed to listen: %v", err)
// }
// lis2, err := net.Listen("tcp", ":3001")
// if err != nil {
//     log.Fatalf("failed to listen: %v", err)
// }

// // Запуск серверов
// go func() {
//     if err := server1.Serve(lis1); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }()
// go func() {
//     if err := server2.Serve(lis2); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }()