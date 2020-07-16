import { has, map, allPass, pipe, propIs } from 'ramda';

export const required = pipe(map(has), allPass);

export const mustBe = (type: any) => pipe(map(propIs(type)), allPass);
