package Database

type MuscleGroup uint8

var Connection *Database = nil

var (
	MUSCLE_GROUP_BACK    MuscleGroup = 0
	MUSCLE_GROUP_BICEP   MuscleGroup = 1
	MUSCLE_GROUP_TRICEP  MuscleGroup = 2
	MUSCLE_GROUP_ABDOMEN MuscleGroup = 3
)
