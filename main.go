package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`
[Intro]
It's called the FENDI Slide, Let's walk through it

[Chorus]
To the left [To the left]
To the right [To the right]
Find your ride and put that whip in drive and do the FENDI Slide
Now show me how you slide [That boy is being hunted
To the left [To the left]
To the right [To the right]
Find your ride, now put that whip in drive
And do the FENDI Slide, now show me how you slide 
Verse 
I'm outside your crib, jump inside
Heard this party jumping, we gon' slide 
Straight to the front of the line
They just happy that we arrived 
I'm the type to finish what I started [Yeah, yeah] 
If I really want it then I bought it [Yeah, yeah]
She so ? when she saw me [She did]
Cuz' I FENDI slide [Slide]
Through the party [I slid]
I don't know your man but he knows me [Ha, ha]
When you gettin' money, you don't get lonely [I am never lonely]
Girl, you working and I really like that [I do]
You bout' to make me miss my flight back
If I text your phone, you better write back [You better]
If one of you is wrong, I cannot write that
I really wish your booty had a Cashapp [Cashapp]

[Chorus]

She sees me do my dance and gets excited
To the left [To the left
To the right [To the right
Find your ride put that whip in drive  
And do the FENDI Slide
Now show me how you slide
To the left [To the left]
To the right, [To the right]
Find your ride, now put that whip in drive
And do the FENDI Slide, show me how you slide
I'm outside your crib, jump inside [Go FENDI, Go FENDI]
Heard this party jumping, we gon' slide [Go FENDI, Go FENDI]
Straight to the front of the line [Go FENDI, Go FENDI]
They just happy that we arrived [Go FENDI]

[Outro]

"Girl you should've been at that party last night, Everybody was doing this dance called the "FENDI Slide" Let me show you how it goes"
Show me how you slide [How you slide]
`)
	})

	var port = os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	app.Listen(":" + port)
} 
