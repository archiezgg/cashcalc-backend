CREATE TABLE "Air_Pricing" (
	"Air_Pricing_ID" serial NOT NULL,
	"Fares" bigint NOT NULL,
	"Doc_Fares" bigint,
	CONSTRAINT "Air_Pricing_pk" PRIMARY KEY ("Air_Pricing_ID")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Road_Pricing" (
	"Road_Pricing" serial NOT NULL,
	"Fares" bigint NOT NULL,
	"Doc_Fares" bigint,
	CONSTRAINT "Road_Pricing_pk" PRIMARY KEY ("Road_Pricing")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Air_Zone" (
	"Air_Zone_ID" int NOT NULL,
	"Name" varchar(255) NOT NULL UNIQUE,
	"Country_ID" int(255) NOT NULL
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Road_Zone" (
	"Road_Zone_ID" int NOT NULL,
	"Name" varchar(255) NOT NULL,
	"Country_ID" int(255) NOT NULL
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Country" (
	"Country_ID" serial NOT NULL,
	"Name" varchar(255) NOT NULL UNIQUE,
	CONSTRAINT "Country_pk" PRIMARY KEY ("Country_ID")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Weight_Air" (
	"Weight_Value" float4 NOT NULL UNIQUE,
	"Air_Zone_ID" int,
	"Air_Pricing_ID" int
) WITH (
  OIDS=FALSE
);





ALTER TABLE "Air_Zone" ADD CONSTRAINT "Air_Zone_fk0" FOREIGN KEY ("Country_ID") REFERENCES "Country"("Country_ID");

ALTER TABLE "Road_Zone" ADD CONSTRAINT "Road_Zone_fk0" FOREIGN KEY ("Country_ID") REFERENCES "Country"("Country_ID");


ALTER TABLE "Weight_Air" ADD CONSTRAINT "Weight_Air_fk0" FOREIGN KEY ("Air_Zone_ID") REFERENCES "Air_Zone"("Air_Zone_ID");
ALTER TABLE "Weight_Air" ADD CONSTRAINT "Weight_Air_fk1" FOREIGN KEY ("Air_Pricing_ID") REFERENCES "Air_Pricing"("Air_Pricing_ID");

