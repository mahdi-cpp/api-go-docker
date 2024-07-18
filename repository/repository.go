package repository

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-docker/model"
	"os/exec"
	"strings"
)

//func Create(temperature model.Temperature) error {
//	err := config.DB.Debug().Create(&temperature).Error
//	return err
//}

func GetDockerImages() []model.Images {

	var images []model.Images
	prg := "docker"
	arg1 := "images"
	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	all := string(stdout)
	lines := strings.Split(strings.ReplaceAll(all, "\r\n", "\n"), "\n")

	for j := 1; j < len(lines)-1; j++ {
		line := strings.Split(lines[j], "   ")
		var strArray [10]string
		var image model.Images
		var i int
		i = 0
		for _, element := range line {
			if len(strings.TrimSpace(element)) > 0 {
				element = strings.TrimSpace(element)
				strArray[i] = element
				i++
			}
		}
		image.Repository = strArray[0]
		image.Tag = strArray[1]
		image.ID = strArray[2]
		image.Created = strArray[3]
		image.Size = strArray[4]

		if strings.Contains(image.Repository, "react") || strings.Contains(image.Repository, "english") {
			image.Type = "react"
		} else if strings.Contains(image.Repository, "golang") || strings.Contains(image.Repository, "tinyhome") {
			image.Type = "golang"
		} else if strings.Contains(image.Repository, "redis") {
			image.Type = "redis"
		} else if strings.Contains(image.Repository, "postgres") {
			image.Type = "postgresql"
		} else if strings.Contains(image.Repository, "rabbitmq") {
			image.Type = "rabbitmq"
		} else if strings.Contains(image.Repository, "nginx") {
			image.Type = "nginx"
		} else if strings.Contains(image.Repository, "node") {
			image.Type = "nodejs"
		} else if strings.Contains(image.Repository, "ubuntu") {
			image.Type = "ubuntu"
		} else {
			image.Type = "container"
		}
		//fmt.Println(image.Image)
		images = append(images, image)
	}
	//fmt.Println(images)

	return images
}

func GetDockerContainers() []model.Container {

	var containers []model.Container
	prg := "docker"
	arg1 := "ps"
	arg2 := "--all"
	cmd := exec.Command(prg, arg1, arg2)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	all := string(stdout)
	lines := strings.Split(strings.ReplaceAll(all, "\r\n", "\n"), "\n")

	for j := 1; j < len(lines)-1; j++ {
		line := strings.Split(lines[j], "   ")
		var strArray [10]string
		var container model.Container
		var i int
		i = 0
		for _, element := range line {
			if len(strings.TrimSpace(element)) > 0 {
				element = strings.TrimSpace(element)
				strArray[i] = element
				i++
			}
		}
		container.ID = strArray[0]
		container.Image = strArray[1]
		container.Command = strArray[2]
		container.Created = strArray[3]
		container.Status = strArray[4]
		if strings.Contains(container.Status, "Exited") {
			container.Name = strArray[5]
			container.Running = false
		} else {
			container.Ports = strArray[5]
			container.Name = strArray[6]
			container.Running = true
		}

		if strings.Contains(container.Name, "Tiny") {
			container.Type = "react"
		} else if strings.Contains(container.Image, "english") {
			container.Type = "golang"
		} else if strings.Contains(container.Image, "redis") {
			container.Type = "redis"
		} else if strings.Contains(container.Image, "postgres") {
			container.Type = "postgresql"
		} else if strings.Contains(container.Image, "rabbitmq") {
			container.Type = "rabbitmq"
		} else if strings.Contains(container.Image, "nginx") {
			container.Type = "nginx"
		} else if strings.Contains(container.Image, "node") {
			container.Type = "nodejs"
		} else {
			container.Type = "container"
		}
		//fmt.Println(container.Image)
		containers = append(containers, container)
	}
	//fmt.Println(containers)

	return containers
}

func GetDockerContainerDetails(id string) model.ContainerDetails {
	fmt.Println("all")

	var containerDetails []model.ContainerDetails
	prg := "docker"
	arg1 := "inspect"
	arg2 := id
	cmd := exec.Command(prg, arg1, arg2)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return model.ContainerDetails{}
	}
	all := string(stdout)
	//fmt.Println(all)

	err = json.Unmarshal([]byte(all), &containerDetails)
	if err != nil {
		fmt.Println(err)
		return model.ContainerDetails{}
	}
	//fmt.Println(containerDetails)

	b, err := json.MarshalIndent(containerDetails, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))

	return containerDetails[0]
}
