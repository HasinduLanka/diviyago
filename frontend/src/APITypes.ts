export interface ConvertedResult {
    Success?: boolean;
    Error?: null;
    TransformedResults?: { [key: string]: TransformedResult };
}

export interface TransformedResult {
    Format?: string;
    VideoCodec?: string;
    AudioCodec?: string;
    Scale?: Scale;
    Quality?: number;
    Data?: string;
    FileExtension?: string
}

export interface Scale {
    Width?: number;
    Height?: number;
}
