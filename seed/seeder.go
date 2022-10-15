package seed

import (
	"log"

	"golang_net_worth_calculator_api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Username: "Gato",
		Email:    "gato@gmail.com",
		Password: "password",
	},
	{
		Username: "Da Bato",
		Email:    "dabato@gmail.com",
		Password: "password",
	},
}

var items = []models.Item{
	{
		Title:       "Mac book pro",
		Price:       1599,
		Description: "Avec la nouvelle puce M2, le MacBook Pro 13 pouces repousse ses propres limites. Fidèle à son design compact, il offre jusqu’à 20 heures d’autonomie1 et se surpasse en matière de performances, tout en gardant la tête froide grâce à son système de refroidissement. Doté d’un sublime écran Retina, d’une caméra FaceTime HD et de micros de qualité studio, il est le plus portable de nos portables pro.",
	},
	{
		Title:       "Iphone 14",
		Price:       1019,
		Description: "Avec iOS 16, vous pouvez personnaliser votre écran verrouillé de façons inédites. Détourez une partie de votre photo pour la mettre en avant. Suivez l’évolution de vos anneaux Activité. Et voyez en direct les informations de vos apps préférées.",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Item{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Item{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Item{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		items[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Item{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed items table: %v", err)
		}
	}
}
