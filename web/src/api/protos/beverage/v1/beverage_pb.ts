// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file beverage/v1/beverage.proto (package beverage.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from enum beverage.v1.BeverageType
 */
export enum BeverageType {
  /**
   * @generated from enum value: BEVERAGE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: BEVERAGE_TYPE_BEER = 1;
   */
  BEER = 1,

  /**
   * @generated from enum value: BEVERAGE_TYPE_COCKTAIL = 2;
   */
  COCKTAIL = 2,

  /**
   * @generated from enum value: BEVERAGE_TYPE_SHOTS = 3;
   */
  SHOTS = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(BeverageType)
proto3.util.setEnumType(BeverageType, "beverage.v1.BeverageType", [
  { no: 0, name: "BEVERAGE_TYPE_UNSPECIFIED" },
  { no: 1, name: "BEVERAGE_TYPE_BEER" },
  { no: 2, name: "BEVERAGE_TYPE_COCKTAIL" },
  { no: 3, name: "BEVERAGE_TYPE_SHOTS" },
]);

/**
 * @generated from enum beverage.v1.Status
 */
export enum Status {
  /**
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: STATUS_INCREASING = 1;
   */
  INCREASING = 1,

  /**
   * @generated from enum value: STATUS_DECREASING = 2;
   */
  DECREASING = 2,

  /**
   * @generated from enum value: STATUS_NO_CHANGE = 3;
   */
  NO_CHANGE = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Status)
proto3.util.setEnumType(Status, "beverage.v1.Status", [
  { no: 0, name: "STATUS_UNSPECIFIED" },
  { no: 1, name: "STATUS_INCREASING" },
  { no: 2, name: "STATUS_DECREASING" },
  { no: 3, name: "STATUS_NO_CHANGE" },
]);

/**
 * @generated from message beverage.v1.Beverage
 */
export class Beverage extends Message<Beverage> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string price = 2;
   */
  price = "";

  /**
   * @generated from field: string base_price = 3;
   */
  basePrice = "";

  /**
   * @generated from field: string name = 4;
   */
  name = "";

  /**
   * @generated from field: int64 percentage_change = 5;
   */
  percentageChange = protoInt64.zero;

  /**
   * @generated from field: beverage.v1.BeverageType type = 6;
   */
  type = BeverageType.UNSPECIFIED;

  /**
   * @generated from field: beverage.v1.Status status = 7;
   */
  status = Status.UNSPECIFIED;

  constructor(data?: PartialMessage<Beverage>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.Beverage";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "price", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "base_price", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "percentage_change", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "type", kind: "enum", T: proto3.getEnumType(BeverageType) },
    { no: 7, name: "status", kind: "enum", T: proto3.getEnumType(Status) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Beverage {
    return new Beverage().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Beverage {
    return new Beverage().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Beverage {
    return new Beverage().fromJsonString(jsonString, options);
  }

  static equals(a: Beverage | PlainMessage<Beverage> | undefined, b: Beverage | PlainMessage<Beverage> | undefined): boolean {
    return proto3.util.equals(Beverage, a, b);
  }
}

/**
 * @generated from message beverage.v1.GetBeveragesRequest
 */
export class GetBeveragesRequest extends Message<GetBeveragesRequest> {
  constructor(data?: PartialMessage<GetBeveragesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.GetBeveragesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBeveragesRequest {
    return new GetBeveragesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBeveragesRequest {
    return new GetBeveragesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBeveragesRequest {
    return new GetBeveragesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetBeveragesRequest | PlainMessage<GetBeveragesRequest> | undefined, b: GetBeveragesRequest | PlainMessage<GetBeveragesRequest> | undefined): boolean {
    return proto3.util.equals(GetBeveragesRequest, a, b);
  }
}

/**
 * @generated from message beverage.v1.GetBeveragesResponse
 */
export class GetBeveragesResponse extends Message<GetBeveragesResponse> {
  /**
   * @generated from field: repeated beverage.v1.Beverage beverages = 1;
   */
  beverages: Beverage[] = [];

  constructor(data?: PartialMessage<GetBeveragesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.GetBeveragesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "beverages", kind: "message", T: Beverage, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBeveragesResponse {
    return new GetBeveragesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBeveragesResponse {
    return new GetBeveragesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBeveragesResponse {
    return new GetBeveragesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetBeveragesResponse | PlainMessage<GetBeveragesResponse> | undefined, b: GetBeveragesResponse | PlainMessage<GetBeveragesResponse> | undefined): boolean {
    return proto3.util.equals(GetBeveragesResponse, a, b);
  }
}

/**
 * @generated from message beverage.v1.CreateBeverageRequest
 */
export class CreateBeverageRequest extends Message<CreateBeverageRequest> {
  /**
   * @generated from field: beverage.v1.Beverage beverage = 1;
   */
  beverage?: Beverage;

  constructor(data?: PartialMessage<CreateBeverageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.CreateBeverageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "beverage", kind: "message", T: Beverage },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBeverageRequest {
    return new CreateBeverageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBeverageRequest {
    return new CreateBeverageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBeverageRequest {
    return new CreateBeverageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateBeverageRequest | PlainMessage<CreateBeverageRequest> | undefined, b: CreateBeverageRequest | PlainMessage<CreateBeverageRequest> | undefined): boolean {
    return proto3.util.equals(CreateBeverageRequest, a, b);
  }
}

/**
 * @generated from message beverage.v1.CreateBeverageResponse
 */
export class CreateBeverageResponse extends Message<CreateBeverageResponse> {
  constructor(data?: PartialMessage<CreateBeverageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.CreateBeverageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBeverageResponse {
    return new CreateBeverageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBeverageResponse {
    return new CreateBeverageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBeverageResponse {
    return new CreateBeverageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateBeverageResponse | PlainMessage<CreateBeverageResponse> | undefined, b: CreateBeverageResponse | PlainMessage<CreateBeverageResponse> | undefined): boolean {
    return proto3.util.equals(CreateBeverageResponse, a, b);
  }
}

/**
 * @generated from message beverage.v1.UpdateBeverageRequest
 */
export class UpdateBeverageRequest extends Message<UpdateBeverageRequest> {
  /**
   * @generated from field: beverage.v1.Beverage beverage = 1;
   */
  beverage?: Beverage;

  constructor(data?: PartialMessage<UpdateBeverageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.UpdateBeverageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "beverage", kind: "message", T: Beverage },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBeverageRequest {
    return new UpdateBeverageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBeverageRequest {
    return new UpdateBeverageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBeverageRequest {
    return new UpdateBeverageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateBeverageRequest | PlainMessage<UpdateBeverageRequest> | undefined, b: UpdateBeverageRequest | PlainMessage<UpdateBeverageRequest> | undefined): boolean {
    return proto3.util.equals(UpdateBeverageRequest, a, b);
  }
}

/**
 * @generated from message beverage.v1.UpdateBeverageResponse
 */
export class UpdateBeverageResponse extends Message<UpdateBeverageResponse> {
  constructor(data?: PartialMessage<UpdateBeverageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.UpdateBeverageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBeverageResponse {
    return new UpdateBeverageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBeverageResponse {
    return new UpdateBeverageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBeverageResponse {
    return new UpdateBeverageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateBeverageResponse | PlainMessage<UpdateBeverageResponse> | undefined, b: UpdateBeverageResponse | PlainMessage<UpdateBeverageResponse> | undefined): boolean {
    return proto3.util.equals(UpdateBeverageResponse, a, b);
  }
}

/**
 * @generated from message beverage.v1.DeleteBeverageRequest
 */
export class DeleteBeverageRequest extends Message<DeleteBeverageRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteBeverageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.DeleteBeverageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteBeverageRequest {
    return new DeleteBeverageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteBeverageRequest {
    return new DeleteBeverageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteBeverageRequest {
    return new DeleteBeverageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteBeverageRequest | PlainMessage<DeleteBeverageRequest> | undefined, b: DeleteBeverageRequest | PlainMessage<DeleteBeverageRequest> | undefined): boolean {
    return proto3.util.equals(DeleteBeverageRequest, a, b);
  }
}

/**
 * @generated from message beverage.v1.DeleteBeverageResponse
 */
export class DeleteBeverageResponse extends Message<DeleteBeverageResponse> {
  constructor(data?: PartialMessage<DeleteBeverageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "beverage.v1.DeleteBeverageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteBeverageResponse {
    return new DeleteBeverageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteBeverageResponse {
    return new DeleteBeverageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteBeverageResponse {
    return new DeleteBeverageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteBeverageResponse | PlainMessage<DeleteBeverageResponse> | undefined, b: DeleteBeverageResponse | PlainMessage<DeleteBeverageResponse> | undefined): boolean {
    return proto3.util.equals(DeleteBeverageResponse, a, b);
  }
}

