import { Controller, Get } from '@nestjs/common';
import { StatusService } from './status.service';

@Controller('status')
export class StatusController {
  constructor(private readonly appService: StatusService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }
}
