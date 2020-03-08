package go_hms_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/timeforaninja/go-hms/structs"
)

type DBhandler interface {
	ConfigGet(key string) string
	ConfigSet(key string, val string)
	ConfigRemove(key string)

	UserGet(userID uint64) *structs.UserSimple
	UserSetByID(userID uint64, user *structs.UserSimple)
	UserRemoveByID(userID uint64)

	GroupGet(groupID uint64) *structs.GroupSimple
	GroupSetByID(groupID uint64, group *structs.GroupSimple)
	GroupRemoveByID(groupID uint64)

	CategorieGet(catID uint64) *structs.CategorieSimple
	CategorieSetByID(catID uint64, cat *structs.CategorieSimple)
	CategorieRemoveByID(catID uint64)

	PlaylistGet(playlistID uint64) *structs.PlaylistSimple
	PlaylistSetByID(playlistID uint64, playlist *structs.PlaylistSimple)
	PlaylistRemoveByID(playlistID uint64)

	VideoGet(videoID uint64) *structs.VideoSimple
	VideoSetByID(videoID uint64, video *structs.VideoSimple)
	VideoRemoveByID(videoID uint64)

	PersonalityGet(persID uint64) *structs.PersonalitySimple
	PersonalitySetByID(persID uint64, pers *structs.PersonalitySimple)
	PersonalityRemoveByID(persID uint64)

	RatingsGet(ratingID uint64) *structs.RatingSimple
	RatingsSetByID(ratingID uint64, rating *structs.RatingSimple)
	RatingsRemoveByID(ratingID uint64)

	CommentsGet(commentID uint64) *structs.CommentSimple
	CommentsSetByID(commentID uint64, comment *structs.CommentSimple)
	CommentsRemoveByID(commentID uint64)

	DownloadQGet(qID uint64) *structs.DownloadQSimple
	DownloadQSetByID(qID uint64, q *structs.DownloadQSimple)
	DownloadQRemoveByID(qID uint64)

	MoviesGet(movieID uint64) *structs.MovieSimple
	MoviesSetByID(movieID uint64, movie *structs.MovieSimple)
	MoviesRemoveByID(movieID uint64)
}

type MYSQLhandler struct {
	db *sql.DB
}

// InitHandler initialises a new mysql handler
func InitMYSQLHandler() (error, DBhandler) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return err, nil
	}
	handler := &MYSQLhandler{db}
	handler.validateTables()
	return nil, handler
}

func (mh MYSQLhandler) validateTables() {
	// TODO: write this
}

func (mh MYSQLhandler) getUser(userID uint64) *structs.User {
	return &structs.User{}
}

func (mh MYSQLhandler) setUser(user *structs.User) {

}
