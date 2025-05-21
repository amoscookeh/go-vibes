package main

import (
	"github.com/amoscookeh/go-vibes"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := vibes.Default()

	// Using vibey HTTP methods

	// VIBE (GET) - get the vibes of all users
	r.VIBE("/users", func(c *vibes.Context) {
		users := []User{
			{ID: "1", Name: "Cosmic Carl", Email: "carl@vibesonly.com"},
			{ID: "2", Name: "Energetic Emma", Email: "emma@highvibrations.org"},
		}
		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"users": users,
			"vibe":  "immaculate",
		})
	})

	// MANIFEST (POST) - manifest a new user into existence
	r.MANIFEST("/users", func(c *vibes.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(vibes.StatusCodes.BadRequest, vibes.Map{
				"message": "Your manifestation energy is misaligned!",
			})
			return
		}

		// Pretend we save to database
		newUser.ID = "3" // assign a new ID

		c.JSON(vibes.StatusCodes.Created, vibes.Map{
			"message": "You have successfully manifested a new user!",
			"user":    newUser,
		})
	})

	// ALIGN (PUT) - align the energy of an existing user
	r.ALIGN("/users/:id", func(c *vibes.Context) {
		userID := c.Param("id")
		var updatedUser User

		if err := c.BindJSON(&updatedUser); err != nil {
			c.JSON(vibes.StatusCodes.BadRequest, vibes.Map{
				"message": "Your alignment energy is disrupted!",
			})
			return
		}

		// Pretend we update in database
		updatedUser.ID = userID

		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"message": "User energy has been realigned!",
			"user":    updatedUser,
		})
	})

	// RELEASE (DELETE) - release a user that no longer serves the collective
	r.RELEASE("/users/:id", func(c *vibes.Context) {
		userID := c.Param("id")

		// Pretend we delete from database

		c.JSON(vibes.StatusCodes.NoContent, vibes.Map{
			"message":     "User has been released back to the universe",
			"released_id": userID,
		})
	})

	r.Run(":8080")
}
