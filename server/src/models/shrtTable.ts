import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model, Sequelize } from "sequelize";

export class Shrt extends Model<InferAttributes<Shrt>, InferCreationAttributes<Shrt>> {
    declare id: CreationOptional<number>;
    declare long_url: string;
    declare slug: string;
    declare visit: CreationOptional<number>;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
}

Shrt.init(
    {
        id: {
            type: DataTypes.INTEGER,
            autoIncrement: true,
            allowNull: false,
            primaryKey: true,
        },
        long_url: {
            type: DataTypes.STRING,
            allowNull: false
        },
        slug: {
            type: DataTypes.STRING,
            allowNull: false
        },
        visit: {
            type: DataTypes.INTEGER,
            allowNull: false,
            defaultValue: 0
        },
        createdAt: DataTypes.DATE,
        updatedAt: DataTypes.DATE,
    },
    {
        sequelize: new Sequelize(`${process.env.DATABASE_URL}`),
        tableName: "Shrts"
    }
);