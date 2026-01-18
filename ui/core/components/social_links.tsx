import tippy from 'tippy.js';

import { Component } from './component';

export class SocialLinks extends Component {
	static buildGitHubLink(): Element {
		const anchor = (
			<a
				href="https://github.com/isfir/wowsims-turtle"
				target="_blank"
				className="github-link link-alt"
				dataset={{ tippyContent: 'Contribute on GitHub' }}>
				<i className="fab fa-github fa-lg" />
			</a>
		);
		tippy(anchor);
		return anchor;
	}

	static buildPatreonLink(): Element {
		const anchor = (
			<a href="https://patreon.com/wowsims" target="_blank" className="patreon-link link-alt" dataset={{ tippyContent: 'Support us on Patreon' }}>
				<i className="fab fa-patreon fa-lg" /> Patreon
			</a>
		);
		tippy(anchor);
		return anchor;
	}
}
