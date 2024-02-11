import { Injectable } from "@nestjs/common";
import { OrdersService } from "./orders.service";
import { OrderStatus } from "./entities/order.entity";

@Injectable()
export class OrderConsumer {
    constructor(
        private orderService: OrdersService) { }

    async consume(msg: { order_id: string, status: OrderStatus}) {
        console.log('Message', msg);

    }
}

class InvalidStatusError extends Error {
    constructor(invalidStatus: string) {
        super(`Invalid status: ${invalidStatus}`);
    }
}