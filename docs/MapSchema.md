Map Schema
==========

Documentation about the standard map description schema in JSON.

## Overall Map Structure

	{
		"name": <String -- Name of map>,
		"authorName": <String -- Name of author>,
		"authorEmail": <String -- Email of author>,
		"width": <Int -- Width of map>,
		"height": <Int -- Height of map>,
		"depth": <Int -- Optional -- 3D depth of map>,
		"configurations": [
			{
				"initial": <JSON Map -- For user to decide>,
				"goal": <JSON Map -- For user to decide>
			}, …
		],
		"obstacles": [
			{
				"dynamic": <Boolean -- Whether obstacle is dynamic>,
				"name": <String -- Optional -- Name of obstacle>,
				"vertices": [
					{
		 				"x": <Int -- X coordinate of vertex>,
						"y": <Int-- Y coordinate of vertex>,
						"z": <Int -- Optional -- Z coordinate of vertex>
					}, …
				]
			}, …
		] 	
	}