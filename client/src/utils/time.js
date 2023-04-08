export function CompareTime(time) {
    const currentTimeUTC0 = new Date().getTime() - new Date().getTimezoneOffset() * 60000;
    const compareTimeUTC0 = new Date(time).getTime() - new Date(time).getTimezoneOffset() * 60000;
    return compareTimeUTC0 > currentTimeUTC0;
}
