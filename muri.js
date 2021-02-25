export { Encode, Decode, Scheme };

const Scheme = "muri://";
function Encode(sa) {
   if(!Array.isArray(sa)){
      throw("invalid array");
   }
    switch (sa.length) {
        case 1:
            return sa[0];
        case 0:
            return "";
    }
    for (let index = 0; index < sa.length; index++) {
        sa[index] = encodeURIComponent(sa[index]);
    }
    return Scheme + encodeURIComponent(sa.join(","));
}

function Decode(s ) {
   
    if (!s.startsWith(Scheme)) {
        // it's actually a URI?
        return [s];
    }
    s = s.slice(Scheme.length);
    const ps = decodeURIComponent(s);
    if (ps == "") {
        throw (`Invalid muri ${s}`);
    }
    let sa = ps.split(",");
    if (sa.length === 0) {
        throw (`Invalid muri ${ps}`);
    }
   for (let index = 0; index < sa.length; index++) {
        sa[index] = decodeURIComponent(sa[index]);
    }
    return sa;
}