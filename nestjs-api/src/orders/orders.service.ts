import { Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Order } from './entities/order.entity';
import { In, Repository } from 'typeorm';
import { Product } from 'src/products/entities/product.entity';

@Injectable()
export class OrdersService {

  constructor(
    @InjectRepository(Order) private orderRepo: Repository<Order>,
    @InjectRepository(Product) private productRepo: Repository<Product>) { }

  async create(createOrderDto: CreateOrderDto) {
    const productsIds = createOrderDto.items.map(item => item.product_id);

    // Esse tratamento é por poder haver mais de um mesmo item no pedido
    // neste caso, será retirados as duplicadas
    const uniqueProductIds = [...new Set(productsIds)]

    const products = await this.productRepo.findBy({
      id: In(productsIds)
    })

    if (products.length !== uniqueProductIds.length) {
      throw new Error(`Algum produto não existe. Produtos passados ${productsIds}, \nprodutos encontrados: ${products.map(p => p.id)}`)
    }

    const order = Order.create({
      client_id: '1',
      items: createOrderDto.items.map(item => {
        const product = products.find(p => p.id === item.product_id)
        return {
          price: product.price,
          product_id: item.product_id,
          quantity: item.quantity
        }
      })
    })

    await this.orderRepo.save(order);
    return order;
  }

  findAll() {
    return this.orderRepo.find();
  }

  findOne(id: number) {
    return `This action returns a #${id} order`;
  }
}


