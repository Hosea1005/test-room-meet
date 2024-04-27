package helper

import "github.com/google/uuid"

func GeneratedUUID() uuid.UUID {
	newUUID := uuid.New()
	return newUUID
}
func ConvertStringToUUID(uuidString string) (uuid.UUID, error) {
	// Mengonversi string ke UUID
	convertedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		// Mengembalikan error jika string tidak dapat dikonversi ke UUID
		return uuid.Nil, err
	}
	// Mengembalikan UUID yang dikonversi dan nil jika berhasil
	return convertedUUID, nil
}
