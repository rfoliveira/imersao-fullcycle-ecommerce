import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { getDataSourceToken } from "@nestjs/typeorm";
import { DataSource } from "typeorm";
// poderia usar outro pacote, mas como já tem esse aqui, usei ele mesmo (tentei...)
// import { UUID } from "typeorm/driver/mongodb/bson.typings";
// outra opção mais padrão porqua a de cima não funciona ou nào sei usar...
import { v4 as uuidv4 } from 'uuid';
// ourta opção
// import { randomUUID } from "crypto";
// ref.: https://developer.mozilla.org/en-US/docs/Web/API/Crypto/randomUUID

// aplicação main fake
async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    await app.init();

    const datasource = app.get<DataSource>(getDataSourceToken());
    // destroi tudo e recria todas as tabelas
    await datasource.synchronize(true);

    // o instrutor usou a criação manualmente, mas preferi fazer dessa forma
    await insertProducts(datasource);
    await app.close();
}

async function insertProducts(ds: DataSource) {
    const productRepo = ds.getRepository('Product')
    const products = []

    for (let index = 1; index <= 10; index++) {
        // Isso não funciona. Fica dizendo que UUID não é uma classe
        // const _id = new UUID(`${index}`).toHexString();
        // console.log(_id);

        products.push({
            id: uuidv4(),
            name: `Product ${index}`,
            description: `Description of the #${index} product`,
            price: index*100,
            image_url: `http://locahost:9000/products/${index}.png`
        })
    }

    await productRepo.insert(products);
}

bootstrap();