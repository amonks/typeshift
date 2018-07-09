/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * DocStringField is a cool field
 */
export type DocStringField = string;
export type NullableStringField = string;
export type OmittedStringField = string;
export type RenamedStringField = string;
export type Int = number;
export type String = string;
export type WeirdlyOmittedStringField = string;
export type PrivateStringField = string;

/**
 * PublicObject is an object
 */
export interface PublicObject {
  DocStringField: DocStringField;
  NullableStringField: NullableStringField | null;
  OmittedStringField: OmittedStringField;
  RenamedStringField: RenamedStringField;
  StructField: StructField;
  WeirdlyOmittedStringField: WeirdlyOmittedStringField;
  privateStringField: PrivateStringField;
  [k: string]: any;
}
export interface StructField {
  Int: Int;
  String: String;
  [k: string]: any;
}
