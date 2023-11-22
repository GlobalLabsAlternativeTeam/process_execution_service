-- CreateEnum
CREATE TYPE "InstanceStatus" AS ENUM ('INSTANCE_STATUS_UNSPECIFIED', 'INSTANCE_STATUS_RUNNING', 'INSTANCE_STATUS_BLOCKED', 'INSTANCE_STATUS_COMPLETED', 'INSTANCE_STATUS_CANCELLED');

-- CreateTable
CREATE TABLE "Instance" (
    "id" SERIAL NOT NULL,
    "doctorId" INTEGER NOT NULL,
    "patientId" INTEGER NOT NULL,
    "status" "InstanceStatus" NOT NULL,
    "schemaInstanceId" TEXT NOT NULL,
    "startedAt" TEXT NOT NULL,
    "finishedAt" TEXT NOT NULL,
    "deletedAt" TEXT NOT NULL,

    CONSTRAINT "Instance_pkey" PRIMARY KEY ("id")
);
