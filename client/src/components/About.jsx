import {useLayoutEffect} from "react";

export const About = () => {
	useLayoutEffect(() => {
		document.title = 'About';
	}, []);
	return (
		<>
			<div>
				<p>This a React client with complied file embed into Golang binary </p>
			</div>
		</>
	)
}