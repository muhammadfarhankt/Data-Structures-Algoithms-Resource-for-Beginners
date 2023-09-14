//Leetcode 1185 Day of the Week
char * dayOfTheWeek(int day, int month, int year){
if(month<3){
        month += 12;
        year--;
    }
    int year_of_century = year % 100;
    int century = year / 100;
    int decade = month;
    int weekday = ( (day + (13 * (decade + 1) / 5) + year_of_century + (year_of_century / 4) + (century / 4) - (2 * century)) % 7);
    char *days[] = {"Saturday","Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"};
    return days[(weekday+7)%7];
}
