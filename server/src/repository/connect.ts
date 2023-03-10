import { Sequelize } from "sequelize";
import { Shrt } from "../models/shrtTable";

export const connectDatabase = async () => {
    let Db: Sequelize
    if (process.env.DATABASE_URL) {
        Db = new Sequelize(process.env.DATABASE_URL)
    } else {
        throw Error("Could't connect to the database")
    }

    // Test connection
    await Db.authenticate();

    // Init
    if (process.env.AUTO_MIGRATE) {
        if (process.env.AUTO_MIGRATE === "1") {
            // Model
            new Shrt()

            // Sync all models
            await Db.sync({ force: true });

            console.log("All models were synchronized successfully.");
        }
    } else {
        throw Error("Could't find the auto migration config")
    }
    console.log('Connection has been established successfully.');

}