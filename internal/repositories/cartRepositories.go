package repositories

import (
	"cart/helper"
	"cart/internal/core/domain"
	"cart/internal/core/ports/cartPort"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type cartInfra struct {
	db *gorm.DB
}

func NewCartRepositories(db *gorm.DB) cartPort.CartRepository {
	return &cartInfra{db}
}

func (r *cartInfra) AddToCart(cart domain.Cart) (domain.Cart, error) {
	newCart := domain.NewCart(cart.ProductId, cart.ProductName, cart.Quantity, cart.Price)
	if err := r.db.Create(newCart).Error; err != nil {
		return domain.Cart{}, helper.PrintErrorMessage(helper.DatabaseError, err.Error())
	}
	return *newCart, nil
}

func (r *cartInfra) DeleteAnItemFromCart(aCart domain.Cart, reference string) (string, error) {
	if err := r.db.Delete(&aCart).Error; err != nil {
		return "domain.Cart{}", helper.PrintErrorMessage(helper.DatabaseError, err.Error())
	}
	return reference, nil
}

func (r *cartInfra) DeleteAllCartItems(carts []domain.Cart) (string, error) {
	if err := r.db.Delete(&carts).Error; err != nil {
		return "domain.Cart{}", helper.PrintErrorMessage(helper.DatabaseError, err.Error())
	}
	return "empty", nil
}

func (r *cartInfra) ReduceQuantity(cart domain.Cart, reference string) (string, error) {
	// get the product price, insert later
	if err := r.db.Model(&cart).Where("reference = ?", reference).Select("Quantity", "Price").Updates(domain.Cart{Quantity: cart.Quantity - 1, Price: cart.Quantity * cart.Price}).Error; err != nil {
		return "domain.Cart{}", helper.PrintErrorMessage(helper.DatabaseError, err.Error())
	}

	return "UPDATE: reduced item quantity", nil
}

func (r *cartInfra) IncreaseQuantity(cart domain.Cart, reference string) (string, error) {
	//add to the quantity of an item
	if err := r.db.Model(&cart).Where("reference = ?", reference).Select("Quantity", "Price").Updates(domain.Cart{Quantity: cart.Quantity + 1, Price: cart.Quantity * cart.Price}).Error; err != nil {
		return "domain.Cart{}", helper.PrintErrorMessage(helper.DatabaseError, err.Error())
	}

	return "UPDATE: increased item quantity", nil
}

type messageRedis struct {
	MessageToSub string `json:"message"`
}

func Idle(db *gorm.DB) (domain.Cart, error) {
	var cart domain.Cart
	var redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	var f func()
	var t *time.Timer
	fmt.Println("redis started")
	if time.Now().Hour() == 16 {
		fmt.Println(db)
		if err := db.Find(&cart).Where("created_at=?", cart.CreatedAt.Hour() >= 72).Error; err != nil {
			return domain.Cart{}, helper.PrintErrorMessage(helper.DatabaseError, err.Error())
		}
		fmt.Println("middle of redis")
		f = func() {
			payload, err := json.Marshal(messageRedis{MessageToSub: "I checked the database if the time is 6 PM and sent a notification about a cart that has been unattended to after 3 days"})
			if err != nil {
				panic(err)
			}

			fmt.Println("publishing to redis")
			if err := redisClient.Publish(context.Background(), "send-idle-notification-message", payload).Err(); err != nil {
				panic(err)
			}
			fmt.Println(messageRedis{})
			t = time.AfterFunc(time.Duration(72)*time.Hour, f)
		}
		t = time.AfterFunc(time.Duration(7)*time.Second, f)
		defer t.Stop()
		fmt.Println("redis ended")
		//simulate doing stuff
		time.Sleep(168 * time.Hour)
	
	}
	return domain.Cart{}, nil
}

