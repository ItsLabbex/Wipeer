package sections

import (
	"Wipe/src/cmd/vars"
	"github.com/gookit/color"
)

func Header() {

	color.Println(`<red>

	 _              _      _                  __          __ _              
	| |            | |    | |                 \ \        / /(_)             
	| |       __ _ | |__  | |__    ___ __  __  \ \  /\  / /  _  _ __    ___ 
	| |      / _´ || '_ \ | '_ \  / _ \\ \/ /   \ \/  \/ /  | || '_ \  / _ \
	| |____ | (_| || |_) || |_) ||  __/ >  <     \  /\  /   | || |_) ||  __/
	|______| \__,_||_.__/ |_.__/  \___|/_/\_\     \/  \/    |_|| .__/  \___|
								   | |
								   |_|
	</>
          	<yellow>------------ By LabbexStudios - v` + vars.Version + ` ------------</>
`)
}
