package video_delrec

import (
	"github.com/spf13/viper"
	"log"
	"os"
)



func Delete(vid int) error{
	err := os.Remove(viper.GetString("app.video_path") + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}

	return nil
}

