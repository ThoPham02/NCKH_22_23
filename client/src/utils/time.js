export function convertTimestampToDateString(timestamp) {
  const date = new Date(timestamp);

  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const day = date.getDate().toString().padStart(2, "0");

  const dateString = `${day}/${month}/${year}`;

  return dateString;
}

export function convertDateToTimestamp(dateString) {
  if (!dateString) return 0;
  const dateParts = dateString.split("/");
  const day = parseInt(dateParts[0], 10);
  const month = parseInt(dateParts[1], 10) - 1;
  const year = parseInt(dateParts[2], 10);

  const date = new Date(year, month, day);
  const timestamp = date.getTime() / 1000;

  return timestamp;
}
