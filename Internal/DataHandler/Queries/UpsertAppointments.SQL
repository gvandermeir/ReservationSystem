BEGIN;

-- Select the row(s) to be updated with a pessimistic lock
SELECT * FROM Appointments WHERE AppointmentID=$1 FOR UPDATE;

-- Perform the update operation
UPDATE Appointments SET PatientID = $2, BookedAt = $3, Confirmed=$4 WHERE AppointmentID=$1;

COMMIT;