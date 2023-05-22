export function convertTimestampToDateString(timestamp) {
  if (timestamp === 0) {
    return "..."
  }
  const date = new Date(timestamp);

  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const day = date.getDate().toString().padStart(2, "0");

  const dateString = `${day}/${month}/${year}`;

  return dateString;
}

export function convertDateToTimestamp(dateString) {
  if (!dateString) return 0;
  var dateObj = new Date(dateString);

  // Lấy timestamp từ đối tượng Date
  var timestamp = dateObj.getTime();

  return timestamp;
}
