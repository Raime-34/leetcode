/**
 * @param {number[]} rec1
 * @param {number[]} rec2
 * @return {boolean}
 */
var isRectangleOverlap = function(rec1, rec2) {
    let flagX = false;
    let flagY = false;

    if (checkVertex(rec1[0], rec1[2], rec2[0], rec2[2]))
        flagX = true;
    if (checkVertex(rec2[0], rec2[2], rec1[0], rec1[2]))
        flagX = true;
    if (checkVertex(rec1[0], rec1[2], rec2[2], rec2[0]))
        flagX = true;
    if (checkVertex(rec2[0], rec2[2], rec1[2], rec1[0]))
        flagX = true;
    if (checkVertex(rec1[1], rec1[3], rec2[1], rec2[3]))
        flagY = true;
    if (checkVertex(rec2[1], rec2[3], rec1[1], rec1[3]))
        flagY = true;
    if (checkVertex(rec1[1], rec1[3], rec2[3], rec2[1]))
        flagY = true;
    if (checkVertex(rec2[1], rec2[3], rec1[3], rec1[1]))
        flagY = true;

    return flagX && flagY;
};

function between(minCoord, maxCoord, currentVertexCoord) {
    if (minCoord < currentVertexCoord && currentVertexCoord < maxCoord) {
        return true;
    }

    return false;
}

function checkVertex(minCoord, maxCoord, currentVertexCoord, otherCurrentVertexCoord) {
    if (minCoord == currentVertexCoord && maxCoord == otherCurrentVertexCoord) {
        return true;
    }
    if (minCoord == otherCurrentVertexCoord && maxCoord == currentVertexCoord) {
        return true;
    }

    if (between(minCoord, maxCoord, currentVertexCoord)) {
        return true;
    }

    if (between(minCoord, maxCoord, otherCurrentVertexCoord)) {
        return true;
    }

    if (
        (minCoord == currentVertexCoord && otherCurrentVertexCoord > maxCoord) ||
        (maxCoord == currentVertexCoord && otherCurrentVertexCoord < minCoord)
    ) {
        return true;
     }

    if (
        (minCoord == otherCurrentVertexCoord && currentVertexCoord > maxCoord) ||
        (maxCoord == otherCurrentVertexCoord && currentVertexCoord < minCoord)
    ) {
        return true;
    }
    
    return false;
}