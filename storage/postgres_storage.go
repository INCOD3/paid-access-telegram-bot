package storage

// import (
// 	"fmt"
//
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )
//
// type PostgresStorage struct {
//   db *gorm.DB
// }
//
// func NewPostgresStorage() (Storage, error) {
//   s := &PostgresStorage{}
//   err := s.initDB()
//   if err != nil {
//     return s, err
//   }
//   return s, nil
// }
//
// func (s *PostgresStorage) getDSN(host, user, password, dbname, port string) string {
//   return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
// }
//
// func (s *PostgresStorage) initDB() error {
//   dsn := s.getDSN("localhost", "postgres", "postgres", "postgres", "5432")
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
//     Logger: logger.Default.LogMode(logger.Silent),
//   })
//   if err != nil {
//     return err
//   }
//   
//   s.db = db
//   return nil
// }
//
