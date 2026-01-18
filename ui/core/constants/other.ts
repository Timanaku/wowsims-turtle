import { Stat } from '../proto/common';

export enum Phase {
	Phase1 = 1,
	Phase2,
	Phase3,
	Phase4,
	Phase5,
	Phase6,
}

export const CURRENT_PHASE = Phase.Phase6;

// Site is served from root
export const BASE_PATH = import.meta.env.BASE_URL;
export const REPO_NAME = BASE_PATH === '/' ? '' : BASE_PATH.replace(/^\//, '').replace(/\/$/, '');

// Get 'elemental_shaman', the pathname part after the repo name
const pathnameParts = window.location.pathname.split('/');
let specDirectory = '';
if (REPO_NAME) {
	const repoPartIdx = pathnameParts.findIndex(part => part == REPO_NAME);
	specDirectory = repoPartIdx == -1 ? '' : pathnameParts[repoPartIdx + 1];
} else {
	// No repo prefix, spec is first non-empty part after leading slash
	specDirectory = pathnameParts[1] || '';
}
export const SPEC_DIRECTORY = specDirectory;

export const GLOBAL_DISPLAY_STATS = [
	Stat.StatHealth,
	Stat.StatStamina,
	Stat.StatArcaneResistance,
	Stat.StatFireResistance,
	Stat.StatNatureResistance,
	Stat.StatFrostResistance,
	Stat.StatShadowResistance,
];

export const GLOBAL_DISPLAY_PSEUDO_STATS = [];

export const GLOBAL_EP_STATS = [];

export enum SortDirection {
	ASC,
	DESC,
}
