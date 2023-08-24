module default {
type User {
        required property first_name -> str;
        required property last_name -> str;
        required property date_of_birth -> datetime;
        required property email -> str {
            constraint exclusive;
        };
        link phone_number -> PhoneNumber;
        required property cv_file_path -> str;
    };

    type PhoneNumber {
        required property number -> str;
        link user -> User;
    };
};