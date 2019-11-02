package route

import (
	. "../controller"
	. "../middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	route.Use(ApiMiddleware(db))
	route.Use(CORS())
	structure := route.Group("/structure")
	{
		structure.GET("/", StructureHome)
		structure.POST("/create", StructureCreate)
		structure.GET("/show/:id", StructureShow)
		structure.PUT("/update/:id", StructureUpdate)
		structure.DELETE("/delete/:id", StructureDelete)
		structure.POST("/delete", StructureDeleteMultiple)
	}

	position := route.Group("/position")
	{
		position.GET("/", PositionHome)
		position.POST("/create", PositionCreate)
		position.GET("/show/:id", PositionShow)
		position.PUT("/update/:id", PositionUpdate)
		position.DELETE("/delete/:id", PositionDelete)
		position.POST("/delete", PositionDeleteMultiple)
	}

	fonction := route.Group("/fonction")
	{
		fonction.GET("/", FonctionHome)
		fonction.POST("/create", FonctionCreate)
		fonction.GET("/show/:id", FonctionShow)
		fonction.PUT("/update/:id", FonctionUpdate)
		fonction.DELETE("/delete/:id", FonctionDelete)
		fonction.POST("/delete", FonctionDeleteMultiple)
	}

	facture := route.Group("/facture")
	{
		facture.GET("/", FactureHome)
		facture.POST("/create", FactureCreate)
		facture.GET("/show/:id", FactureShow)
		facture.PUT("/update/:id", FactureUpdate)
		facture.DELETE("/delete/:id", FactureDelete)
		facture.POST("/delete", FactureDeleteMultiple)
	}

	categorie_article := route.Group("/categorie_article")
	{
		categorie_article.GET("/", CategorieArticleHome)
		categorie_article.POST("/create", CategorieArticleCreate)
		categorie_article.GET("/show/:id", CategorieArticleShow)
		categorie_article.PUT("/update/:id", CategorieArticleUpdate)
		categorie_article.DELETE("/delete/:id", CategorieArticleDelete)
		categorie_article.POST("/delete", CategorieArticleDeleteMultiple)
	}

	bon_livraison := route.Group("/bon_livraison")
	{
		bon_livraison.GET("/", BonLivraisonHome)
		bon_livraison.POST("/create", BonLivraisonCreate)
		bon_livraison.GET("/show/:id", BonLivraisonShow)
		bon_livraison.PUT("/update/:id", BonLivraisonUpdate)
		bon_livraison.DELETE("/delete/:id", BonLivraisonDelete)
		bon_livraison.POST("/delete", BonLivraisonDeleteMultiple)
	}

	bon_commande := route.Group("/bon_commande")
	{
		bon_commande.GET("/", BonCommandeHome)
		bon_commande.POST("/create", BonCommandeCreate)
		bon_commande.GET("/show/:id", BonCommandeShow)
		bon_commande.PUT("/update/:id", BonCommandeUpdate)
		bon_commande.DELETE("/delete/:id", BonCommandeDelete)
		bon_commande.POST("/delete", BonCommandeDeleteMultiple)
	}

	acteur := route.Group("/acteur")
	{
		acteur.GET("/", ActeurHome)
		acteur.POST("/create", ActeurCreate)
		acteur.GET("/show/:id", ActeurShow)
		acteur.PUT("/update/:id", ActeurUpdate)
		acteur.DELETE("/delete/:id", ActeurDelete)
		acteur.POST("/delete", ActeurDeleteMultiple)
	}

	abonnement := route.Group("/abonnement")
	{
		abonnement.GET("/", AbonnementHome)
		abonnement.POST("/create", AbonnementCreate)
		abonnement.GET("/show/:id", AbonnementShow)
		abonnement.PUT("/update/:id", AbonnementUpdate)
		abonnement.DELETE("/delete/:id", AbonnementDelete)
		abonnement.POST("/delete", AbonnementDeleteMultiple)
	}

	article := route.Group("/article")
	{
		article.GET("/", ArticleHome)
		article.POST("/create", ArticleCreate)
		article.GET("/show/:id", ArticleShow)
		article.PUT("/update/:id", ArticleUpdate)
		article.DELETE("/delete/:id", ArticleDelete)
		article.POST("/delete", ArticleDeleteMultiple)
	}

	return route
}
