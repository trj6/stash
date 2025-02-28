package models

type GalleryFilterType struct {
	And     *GalleryFilterType    `json:"AND"`
	Or      *GalleryFilterType    `json:"OR"`
	Not     *GalleryFilterType    `json:"NOT"`
	ID      *IntCriterionInput    `json:"id"`
	Title   *StringCriterionInput `json:"title"`
	Details *StringCriterionInput `json:"details"`
	// Filter by file checksum
	Checksum *StringCriterionInput `json:"checksum"`
	// Filter by path
	Path *StringCriterionInput `json:"path"`
	// Filter by zip file count
	FileCount *IntCriterionInput `json:"file_count"`
	// Filter to only include galleries missing this property
	IsMissing *string `json:"is_missing"`
	// Filter to include/exclude galleries that were created from zip
	IsZip *bool `json:"is_zip"`
	// Filter by rating expressed as 1-5
	Rating *IntCriterionInput `json:"rating"`
	// Filter by rating expressed as 1-100
	Rating100 *IntCriterionInput `json:"rating100"`
	// Filter by organized
	Organized *bool `json:"organized"`
	// Filter by average image resolution
	AverageResolution *ResolutionCriterionInput `json:"average_resolution"`
	// Filter to only include scenes which have chapters. `true` or `false`
	HasChapters *string `json:"has_chapters"`
	// Filter to only include galleries with this studio
	Studios *HierarchicalMultiCriterionInput `json:"studios"`
	// Filter to only include galleries with these tags
	Tags *HierarchicalMultiCriterionInput `json:"tags"`
	// Filter by tag count
	TagCount *IntCriterionInput `json:"tag_count"`
	// Filter to only include galleries with performers with these tags
	PerformerTags *HierarchicalMultiCriterionInput `json:"performer_tags"`
	// Filter to only include galleries with these performers
	Performers *MultiCriterionInput `json:"performers"`
	// Filter by performer count
	PerformerCount *IntCriterionInput `json:"performer_count"`
	// Filter galleries that have performers that have been favorited
	PerformerFavorite *bool `json:"performer_favorite"`
	// Filter galleries by performer age at time of gallery
	PerformerAge *IntCriterionInput `json:"performer_age"`
	// Filter by number of images in this gallery
	ImageCount *IntCriterionInput `json:"image_count"`
	// Filter by url
	URL *StringCriterionInput `json:"url"`
	// Filter by date
	Date *DateCriterionInput `json:"date"`
	// Filter by created at
	CreatedAt *TimestampCriterionInput `json:"created_at"`
	// Filter by updated at
	UpdatedAt *TimestampCriterionInput `json:"updated_at"`
}

type GalleryUpdateInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	ID               string   `json:"id"`
	Title            *string  `json:"title"`
	URL              *string  `json:"url"`
	Date             *string  `json:"date"`
	Details          *string  `json:"details"`
	Rating           *int     `json:"rating"`
	Rating100        *int     `json:"rating100"`
	Organized        *bool    `json:"organized"`
	SceneIds         []string `json:"scene_ids"`
	StudioID         *string  `json:"studio_id"`
	TagIds           []string `json:"tag_ids"`
	PerformerIds     []string `json:"performer_ids"`
	PrimaryFileID    *string  `json:"primary_file_id"`
}

type GalleryDestroyInput struct {
	Ids []string `json:"ids"`
	// If true, then the zip file will be deleted if the gallery is zip-file-based.
	// If gallery is folder-based, then any files not associated with other
	// galleries will be deleted, along with the folder, if it is not empty.
	DeleteFile      *bool `json:"delete_file"`
	DeleteGenerated *bool `json:"delete_generated"`
}
