package project

import (
	"day5/project-api/database"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBProject struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	DCode       string
	OwnerName   string
	PCode       string
	ProjectName string
	Status      string
}

var db = database.Connect()

func init() {
	db.AutoMigrate(&DBProject{})
}

type Service struct{}

func (s *Service) ListProjects(ctx *gin.Context) {
	var rawProjects []DBProject
	res := db.Find(&rawProjects)
	if res.RowsAffected > 0 {
		projects := make([]map[string]any, len(rawProjects))
		for r := 0; r < len(rawProjects); r++ {
			projects[r] = gin.H{
				"dCode":       rawProjects[r].DCode,
				"pCode":       rawProjects[r].PCode,
				"status":      rawProjects[r].Status,
				"ownerName":   rawProjects[r].OwnerName,
				"projectName": rawProjects[r].ProjectName,
			}
		}
		ctx.JSON(200, projects)
	} else {
		ctx.JSON(404, gin.H{
			"message": "Project not found",
		})
		return
	}
}

func (s *Service) CreateProject(ctx *gin.Context) {
	var createProject CreateProject
	_ = ctx.Bind(&createProject)

	query := &DBProject{
		DCode:       createProject.DCode,
		OwnerName:   createProject.OwnerName,
		ProjectName: createProject.ProjectName,
		Status:      "active",
	}

	if res := db.Create(query); res.RowsAffected > 0 {
		ctx.JSON(200, gin.H{
			"dCode":       createProject.DCode,
			"pCode":       fmt.Sprintf("P%03d", query.ID),
			"status":      query.Status,
			"ownerName":   createProject.OwnerName,
			"projectName": createProject.ProjectName,
		})
	} else {
		ctx.JSON(400, gin.H{"msg": "Project not found"})
		return
	}

	ctx.JSON(201, Project{
		DCode:       query.DCode,
		PCode:       fmt.Sprintf("P%03d", query.ID),
		ProjectName: query.ProjectName,
		OwnerName:   &query.OwnerName,
		Status:      ProjectStatus(query.Status),
	})
}

func (s *Service) GetProject(ctx *gin.Context, code string) {
	var rawProject DBProject

	wop := strings.TrimPrefix(code, "P")
	fmt.Println(wop)
	c, _ := strconv.Atoi(wop)

	res := db.First(&rawProject, "id = ?", c)
	if res.RowsAffected > 0 {
		//projects := make([]map[string]any, len(rawProjects))
		ctx.JSON(200, gin.H{
			"dCode":       rawProject.DCode,
			"pCode":       rawProject.PCode,
			"status":      rawProject.Status,
			"ownerName":   rawProject.OwnerName,
			"projectName": rawProject.ProjectName,
		})
	} else {
		ctx.JSON(404, gin.H{
			"message": "Project not found",
		})
		return
	}
}

func (s *Service) UpdateProject(c *gin.Context, code string) {

}
