import { Column, CreateDateColumn, Entity, OneToMany, PrimaryGeneratedColumn } from "typeorm";
import { OrderItem } from "./order-item.entity";

export enum OrderStatus {
    PENDING = 'pending',
    PAID = 'paid',
    FAILED = 'failed'
}

export type CreateOrderCommand = {
    client_id: string;
    items: {
        product_id: string;
        quantity: number;
        price: number;
    }[];
}

@Entity()
export class Order {
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Column()
    total: number;

    @Column({ type: 'decimal', precision: 10, scale: 2})
    client_id: string;

    @Column()
    status: OrderStatus = OrderStatus.PENDING;

    @CreateDateColumn()
    created_at: Date;

    @OneToMany(() => OrderItem, (item) => item.order)
    items: OrderItem[];

    static create(input: CreateOrderCommand) {
        const order = new Order();
        order.client_id = input.client_id;
        order.items = input.items.map(item => {
            const orderItem = new OrderItem();

            orderItem.product_id = item.product_id;
            orderItem.quantity = item.quantity;
            orderItem.price = item.price;

            return orderItem;
        });
        order.total = order.items.reduce((sum, item) => {
            return sum + item.price + item.quantity;
        }, 0);

        return order;
    }

    pay() {
        if (this.status === OrderStatus.PAID) {
          throw new Error('Order already paid');
        }
    
        if (this.status === OrderStatus.FAILED) {
          throw new Error('Order already failed');
        }
    
        this.status = OrderStatus.PAID;
    }

    fail() {
        if (this.status === OrderStatus.FAILED) {
          throw new Error('Order already failed');
        }
    
        if (this.status === OrderStatus.PAID) {
          throw new Error('Order already paid');
        }
    
        this.status = OrderStatus.FAILED;
    }
}