import gsap from 'gsap';
import { MotionPathPlugin } from 'gsap/MotionPathPlugin';
import { PixiPlugin } from 'gsap/PixiPlugin';

export function init() {
  gsap.registerPlugin(MotionPathPlugin, PixiPlugin);
}
