CREATE MIGRATION m1algyjyiiywb42fijdwupumn7snxzni5om3ur5laydevqsioch3rq
    ONTO initial
{
  CREATE TYPE default::PhoneNumber {
      CREATE REQUIRED PROPERTY number: std::str;
  };
  CREATE TYPE default::User {
      CREATE LINK phone_number: default::PhoneNumber;
      CREATE REQUIRED PROPERTY cv_file_path: std::str;
      CREATE REQUIRED PROPERTY date_of_birth: std::datetime;
      CREATE REQUIRED PROPERTY email: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY first_name: std::str;
      CREATE REQUIRED PROPERTY last_name: std::str;
  };
  ALTER TYPE default::PhoneNumber {
      CREATE LINK user: default::User;
  };
};
