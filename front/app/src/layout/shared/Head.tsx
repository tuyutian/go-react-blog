type Props = {
    title: string;
    description?: string;
};

const SERVICE_NAME = import.meta.env.VITE_SERVICE_NAME

export const Head = ({title, description}: Props) => {
    return (
        <>
            <title>{`${title} | ${SERVICE_NAME}`}</title>
            <meta name="description" content={description ?? `- ${SERVICE_NAME}`}/>
            <meta property="og:title" content={`${title} | ${SERVICE_NAME}`}/>
            <meta property="og:description" content={description ?? `- ${SERVICE_NAME}`}/>
            <meta name="robots" content="noindex"/>
        </>
    )
};
