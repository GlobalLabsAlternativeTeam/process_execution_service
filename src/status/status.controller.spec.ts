import { Test, TestingModule } from '@nestjs/testing';
import { StatusController } from './status.controller';
import { StatusService } from './status.service';

describe('StatusController', () => {
  let statusController: StatusController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [StatusController],
      providers: [StatusService],
    }).compile();

    statusController = app.get<StatusController>(StatusController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(statusController.getHello()).toBe('Hello World!');
    });
  });
});
