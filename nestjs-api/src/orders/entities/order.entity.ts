import { Product } from "src/products/entities/product.entity";
import { Column, CreateDateColumn, JoinColumn, ManyToOne, PrimaryGeneratedColumn } from "typeorm";

export enum OrderStatus {
    PENDING = 'pending',
    PAID = 'paid',
    FAILED = 'failed'
}

export class Order {
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Column()
    total: number;

    @Column()
    client_id: string;

    @Column()
    status: OrderStatus = OrderStatus.PENDING;

    @CreateDateColumn()
    created_at: Date;
}

export class OrderItem {
    @PrimaryGeneratedColumn()
    id: number;

    @Column({type: 'int'})
    quantity: number;
    
    @Column()
    price: number;

    @ManyToOne(() => Product)
    @JoinColumn({ name: 'product_id' }) // padrão productId
    product: Product;

    @Column()
    product_id: string;

    @ManyToOne(() => Order)
    order: Order;    
}
