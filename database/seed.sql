-- insert into users
Insert INTO Users (id,userName, password, isAdmin) 
VALUES 
	(1,"tom","123",false),
	(2,"sally","123",false),
	(3,"isaac","123",true)
; 

Insert INTO Notifications (id,userID,message,isRead)
VALUES
	(1,1,"welcome to the app",false),
	(2,2,"welcome to the app",false),
	(3,3,"welcome to the app",true),
	(4,1,"remember to workout today!",true),
	(5,2,"remember to workout today!",false),
	(6,3,"remember to workout today!",true)
;

Insert INTO Activities (id, name, userID, duration, intensity) VALUES
	(1,"Running",1,20,"Medium"),
	(2,"Swimming",1,20,"High"),
	(3,"Lifting",2,20,"Medium"),
	(4,"Hiking",2,20,"Low"),
	(5,"Basketball",3,20,"Medium"),
	(6,"Running",3,20,"High")
;
