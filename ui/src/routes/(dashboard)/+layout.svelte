<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import {
		faHome,
		faBolt,
		faCreditCard,
		faDisplay,
		faCloud,
		faFilePen,
		faRightFromBracket,
	} from '@fortawesome/free-solid-svg-icons';
	import NavItem from '$lib/navigation/NavItem.svelte';
</script>

<div class="drawer drawer-end drawer-mobile h-full p-4">
	<input id="side-menu" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content card bg-base-200 shadow-xl p-4">
		<slot />
	</div>
	<div class="drawer-side card bg-base-200 shadow-xl ml-4">
		<label for="side-menu" class="drawer-overlay" />
		<ul class="menu p-4 w-64 text-base-content">
			<li>
				<NavItem href="/">
					<Fa icon={faHome} />
					<span>Dashboard</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/evse">
					<Fa icon={faBolt} />
					<span>EVSE</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/card-reader">
					<Fa icon={faCreditCard} />
					<span>Card Reader</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/display">
					<Fa icon={faDisplay} />
					<span>Display</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/ocpp">
					<Fa icon={faCloud} />
					<span>OCPP</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/logs">
					<Fa icon={faFilePen} />
					<span>Logs</span>
				</NavItem>
			</li>
			<li>
				<NavItem href="/login">
					<Fa icon={faRightFromBracket} />
					<span>Logout</span>
				</NavItem>
			</li>
		</ul>
	</div>
</div>
