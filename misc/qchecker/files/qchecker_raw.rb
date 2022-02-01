## cat qchecker_raw.rb | sed -e s/##.*$// | ruby > qchecker.rb

## Flag flag is correct if
## f = bytes_to_long(flag[::-1]),
## f % 0x1000000000000000d == 0,
## f % 0x10000000000000025 == 0,
## f % 0x10000000000000033 == 0 and
## f % 0x10000000000000051 == 0.

eval$uate=%w(
  ## a: this program
  a = %(eval$uate=%w(#{$uate})*"");

  ## B???: padding
  Bxxx =

  ## b: string -> [int]
  b = -> a {
    a.split(?+).map{|b|
      b.to_i(36)
    }
  };

  ## c: chracters
  ## ["C", "E", "G", "N", "O", "R", "S", "T", "W", "."]
  c = b["awyiv4fjfkuu2pkv+awyiv4fvut716g3j+axce5e4pxrogszr3+5i0omfd5dm9xf9q7+axce5e4khrz21ypr+5htqqi9iasvmjri7+axcc76i03zrn7gu7+cbt4m8xybr3cb27+1ge6snjex10w3si9+1k8vdb4fzcys2yo0"];

  ## Current state
  ## d: phase (0: input, 1: wrong, 2: correct)
  ## e: length of current flag
  ## f: flag mod 0x1000000000000000d
  ## g: flag mod 0x10000000000000025
  ## h: flag mod 0x10000000000000033
  ## i: flag mod 0x10000000000000051
  d, e, f, g, h, i = b["0+0+zeexaxq012eg+k2htkr1olaj6+3cbp5mnkzllt3+2qpvamo605t7j"];

  ## Update state

  ## j = ARGV[0]
  ## if ARGV[0] && d==0 then
  (j = eval(?A<<82<<71<<86)[0]) && d==0 && (
    e += 1;

    k = 2**64;

    ## l(x, m): (x-j)/256 mod m
    l = -> (a, b) {
      (a-j.ord) * 256.pow(b-2, b) % b
    };

    f = l[f, k+13];
    g = l[g, k+37];
    h = l[h, k+51];
    i = l[i, k+81];

    ## Flag is correct if e==32 and [f, g, h, i]==[0, 0, 0, 0]
    j==?} && (
      d = e==32 && f+g+h+i==0 ? 2 : 1
    );

    a.sub!(/"0.*?"/,
      '"0' + [d, e, f, g, h, i].map{|x| x.to_s(36)}*?+ <<34)
  );

  ## Format program
  srand(f);

  ## k: ["SECCON", "WRONG.", "CORRECT"]
  k = b["7acw+jsjm+46d84"];
  ## l: k[d].length
  l = d==2 ? 7 : 6;
  ## m: template
  m = [?#*(l*20)<<10]*11*"";

  l.times{|a|
    ## rotate b*90 degree
    b = d==0 && e!=0 ? rand(4) : 0;
    ## e: y
    9.times{|e|
      ## f: x
      9.times{|f|
       (c[k[d]/10**a%10]>>(e*9+f)&1)!=0 && (
         ## g = x, h = y
         g = f;
         h = e;
         b.times{
           g, h = h, 8-g
         };

         ## t = (h+1)*(l*20+1)+a*20+g*2+1;
         t = (h*l+l+a)*20+h+g*2+2;
         m[t] = m[t+1] = ""<<32
       )
      }
    }
  };

  ## Remove paddings
  a.sub!(/B.*?=/, "B=");
  ## n: padding length
  n = m.count(?#)-a.length;
  a.sub!("B=", "B#{(1..n).map{(rand(26)+97).chr}*""}=");

  ## fill m with a
  o = 0;
  m.length.times{|b|
    m[b]==?# && o<a.length && (
      m[b] = a[o];
      o += 1
    )
  };

  puts(m)
)*""
