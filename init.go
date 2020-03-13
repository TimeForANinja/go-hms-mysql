package go_hms_mysql

import (
	"database/sql"
	"fmt"
)

// InitTables creates default tables in a database
func InitTables(db *sql.DB) {
	querys := []string{
		sqlQueries.QueryCreateConfig, // config.go

		sqlQueries.QueryCreateCategories, // categories.go

		sqlQueries.QueryCreateLinkCategorieUserPermission,  // categories.go
		sqlQueries.QueryCreateLinkCategoriePersonality,     // categories.go
		sqlQueries.QueryCreateLinkCategorieVideo,           // categories.go
		sqlQueries.QueryCreateLinkCategorieMovie,           // categories.go
		sqlQueries.QueryCreateLinkCategorieGroupPermission, // categories.go
		sqlQueries.QueryCreateLinkCategoriePlaylist,        // categories.go

		sqlQueries.QueryCreateUsers,                // users.go
		sqlQueries.QueryCreateGroups,               // groups.go
		sqlQueries.QueryCreateMovies,               // movies.go
		sqlQueries.QueryCreateVideos,               // videos.go
		sqlQueries.QueryCreateDownloadQ,            // downloadQ.go
		sqlQueries.QueryCreatePersonalities,        // personalities.go
		sqlQueries.QueryCreateLinkPersonalityVideo, // personalities.go
		sqlQueries.QueryCreateLinkPersonalityMovie, // personalities.go
		sqlQueries.QueryCreatePlaylists,            // playlists.go
		sqlQueries.QueryCreateLinkUserGroup,        // users.go
		sqlQueries.QueryCreateComments,             // users.go
		sqlQueries.QueryCreateRatings,              // users.go
		sqlQueries.QueryCreateLinkIdenticalVideos,  // videos.go
		sqlQueries.QueryCreateLinkVideoPlaylist,    // videos.go

		sqlQueries.QueryDefaultAdminUser,       // defaults.go
		sqlQueries.QueryDefaultEveryoneGroup,   // defaults.go
		sqlQueries.QueryDefaultGroup,           // defaults.go
		sqlQueries.QueryDefaultAdminUserGroups, // defaults.go
	}
	for i := 0; i < len(querys); i++ {
		query := querys[i]
		_, err := db.Exec(query)
		if err != nil {
			panic(fmt.Errorf("Error \"%s\" running query:\"%s\"", err, query))
		}
	}
}