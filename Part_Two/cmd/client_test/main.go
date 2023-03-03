package main

import (
	"log"
	"math/rand"
	"time"
	
	pb "github.com/Krokin/gocloudcamp/Part_Two/proto"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewPlaylistClient(conn)

	log.Println("Запускаем пустой плейлист")
	if _, err := client.Play(context.TODO(), &empty.Empty{}); err != nil {
		log.Println(err)
	}
	log.Println("Добавляем 5 песен в плейлист")
	AddSong(5, client)
	log.Println("Запускаем плейлист")
	if _, err := client.Play(context.TODO(), &empty.Empty{}); err != nil {
		log.Println(err)
	}

	log.Println("Получаем играющую песню")
	res, err := client.ReadSong(context.TODO(), &empty.Empty{})
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	log.Println("Пытаемся удалить играющую песню")
	if _, err := client.Delete(
		context.TODO(),
		&pb.DeleteSongRequest{Author: res.GetSong().GetAuthor(), Title: res.GetSong().GetTitle()});
		err != nil {
		log.Println(err)
	}
	log.Println("Включаем следующий трек")
	if _, err := client.Next(context.TODO(), &empty.Empty{}); err != nil {
		log.Println(err)
	}
	log.Println("Пытаемся удалить тот же уже неиграющий трек")
	_, err = client.Delete(
		context.TODO(),
		&pb.DeleteSongRequest{Author: res.GetSong().GetAuthor(), Title: res.GetSong().GetTitle()})
	if err != nil {
		log.Println(err)
	}
	log.Println("Читаем плейлист")
	playlist, err := client.ReadPlaylist(context.TODO(), &empty.Empty{})
	if err != nil {
		log.Println(err)
	}
	log.Println(playlist)
	log.Println("Выключаем сервер")
}

func AddSong(n int, cl pb.PlaylistClient) {
	type song struct {
		a string
		t string
		d time.Duration
	}
	songs := []song{}
	for i := n; i > 0; i-- {
		songs = append(songs, song{gofakeit.Name(), gofakeit.HipsterWord(), time.Second * time.Duration(rand.Intn(150))})
	}
	for _, s := range songs {
		_, err := cl.Create(
			context.TODO(),
			&pb.CreateSongRequest{
				Song: &pb.Song{
					Author: s.a,
					Title:  s.t,
					Dur:    durationpb.New(s.d),
				},
			})
		if err != nil {
			log.Println(err)
		}
	}
}
