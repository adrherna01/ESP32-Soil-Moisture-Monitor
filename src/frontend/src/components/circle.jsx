import React from 'react';

export default function Circle({
  size = 100,
  color = 'blue',
  visiblePercentage = 100,
  style = {},
}) {
  const clipPath = getClipPathFromPercentage(visiblePercentage);

  return (
    <div
      style={{
        width: size,
        height: size,
        backgroundColor: color,
        borderRadius: '50%',
        clipPath,
        ...style,
      }}
    />
  );
}

function getClipPathFromPercentage(pct) {
  const clamped = Math.max(0, Math.min(100, pct));
  const rightClip = 100 - clamped;
  return `inset(0% ${rightClip}% 0% 0%)`; // shows left pct%
}