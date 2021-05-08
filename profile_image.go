package faker

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const profileImageBaseUrl = "https://thispersondoesnotexist.com/image"

// ProfileImage  is a faker struct for ProfileImage
type ProfileImage struct {
	faker *Faker
}

// Image generates a *os.File with a random profile image using the thispersondoesnotexist.com service
func (pi ProfileImage) Image() *os.File {
	resp, err := get(profileImageBaseUrl)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error while requesting", profileImageBaseUrl, ":", err)
	}

	f, err := ioutil.TempFile(os.TempDir(), "profil-picture-img-*.jfif")
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)
	return f
}
