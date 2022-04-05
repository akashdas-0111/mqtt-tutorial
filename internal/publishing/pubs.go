package publishing

import (
	"bufio"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func publish(client mqtt.Client,topic string) {
	scanner := bufio.NewScanner(os.Stdin)
	for{
		scanner.Scan()
		text :=scanner.Text()
		if(text=="exit"){
			break
		}
        token := client.Publish(topic, 1, false, text)
        token.Wait()
	}

}