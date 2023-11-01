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
		projects := make([]Project, len(rawProjects))
		for r := 0; r < len(rawProjects); r++ {
			ownerName := rawProjects[r].OwnerName
			projects[r] = Project{
				DCode:       rawProjects[r].DCode,
				PCode:       fmt.Sprintf("P%03d", rawProjects[r].ID),
				Status:      ProjectStatus(rawProjects[r].Status),
				OwnerName:   &ownerName,
				ProjectName: rawProjects[r].ProjectName,
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
		ctx.JSON(201, Project{
			DCode:       query.DCode,
			PCode:       fmt.Sprintf("P%03d", query.ID),
			ProjectName: query.ProjectName,
			OwnerName:   &query.OwnerName,
			Status:      ProjectStatus(query.Status),
		})

	} else {
		ctx.JSON(400, gin.H{"msg": "Project not found"})
		return
	}
}

func (s *Service) GetProject(ctx *gin.Context, code string) {
	var rawProject DBProject

	wop := strings.TrimPrefix(code, "P")
	fmt.Println(wop)
	c, _ := strconv.Atoi(wop)

	res := db.First(&rawProject, "id = ?", c)
	if res.RowsAffected > 0 {
		//projects := make([]map[string]any, len(rawProjects))
		ctx.JSON(200, Project{
			DCode:       rawProject.DCode,
			PCode:       fmt.Sprintf("P%03d", rawProject.ID),
			ProjectName: rawProject.ProjectName,
			OwnerName:   &rawProject.OwnerName,
			Status:      ProjectStatus(rawProject.Status),
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
