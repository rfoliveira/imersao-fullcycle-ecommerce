import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ProductsModule } from './products/products.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { hostname } from 'os';
import { Product } from './products/entities/product.entity';
import { OrdersModule } from './orders/orders.module';

@Module({
  imports: [
    // forRoot apenas para o módulo principal (pai)
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'localhost',
      port: 3306,
      username: 'root',
      password: 'root',
      database: 'nest',
      entities: [
        Product
      ],
      synchronize: true,  // sincroniza o código com o banco
      logging: true // exibe o log no console
    }),
    ProductsModule,
    OrdersModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
