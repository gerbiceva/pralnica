package paths

import (
	"encoding/json"
	"net/http"
	dbclient "rso/m/dbClient"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetUserReservations retrieves a user's reservations.
func GetUserReservations(w http.ResponseWriter, r *http.Request) {
	userUUID := chi.URLParam(r, "uuid")
	var pgUUID pgtype.UUID
	parsedUUID, err := uuid.Parse(userUUID)
	if err != nil {
		http.Error(w, "Invalid UUID string: %v", http.StatusInternalServerError)
	}
	pgUUID.Bytes = parsedUUID
	pgUUID.Valid = true

	reservations, err := dbclient.Client.GetUserReservations(r.Context(), pgtype.UUID{
		Bytes: [16]byte(parsedUUID),
		Valid: false,
	})

	if err != nil {
		http.Error(w, "Unable to fetch reservations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

// GetReservationsByMonth retrieves all reservations for a specific month.
func GetReservationsByMonth(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")
	var pgMonth pgtype.Date
	if err := pgMonth.Scan(month); err != nil {
		http.Error(w, "Invalid month format", http.StatusBadRequest)
		return
	}

	reservations, err := dbclient.Client.GetReservationsByMonth(r.Context(), pgMonth)
	if err != nil {
		http.Error(w, "Unable to fetch reservations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

// DeleteReservation deletes a specific reservation.
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	reservationID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(reservationID)
	if err != nil {
		http.Error(w, "Invalid reservation ID. Not a number", http.StatusInternalServerError)
		return
	}

	err = dbclient.Client.DeleteReservation(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Unable to delete reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
