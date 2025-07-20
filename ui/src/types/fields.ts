export enum FieldType {
  STRING = 'string',
  TEXT = 'text',
  NUMBER = 'number',
  BOOLEAN = 'boolean',
  DATE = 'date',
  IMAGE = 'image',
  URL = 'url',
}


export const VALID_FIELD_TYPES = Object.values(FieldType);


export interface Field {
  name: string;
  type: FieldType;
}