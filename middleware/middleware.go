package middleware

import (
	. "../repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ApiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		r := make(map[string]interface{})

		r["Structure"] = NewStructureRepository(db)
		r["Position"] = NewPositionRepository(db)
		r["Abonnement"] = NewAbonnementRepository(db)
		r["Acteur"] = NewActeurRepository(db)
		r["Article"] = NewArticleRepository(db)
		r["BonCommande"] = NewBonCommandeRepository(db)
		r["BonLivraison"] = NewBonLivraisonRepository(db)
		r["CategorieArticle"] = NewCategorieArticleRepository(db)
		r["Facture"] = NewFactureRepository(db)
		r["Fonction"] = NewFonctionRepository(db)
		r["Livraison"] = NewLivraisonRepository(db)

		c.Set("structrepo", r["Structure"])
		c.Set("positionrepo", r["Position"])
		c.Set("abonnementrepo", r["Abonnement"])
		c.Set("acteurrepo", r["Acteur"])
		c.Set("articlerepo", r["Article"])
		c.Set("boncommanderepo", r["BonCommande"])
		c.Set("bonlivraisonrepo", r["BonLivraison"])
		c.Set("categoriearticlerepo", r["CategorieArticle"])
		c.Set("facturerepo", r["Facture"])
		c.Set("fonctionrepo", r["Fonction"])
		c.Set("livraisonrepo", r["Livraison"])

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
